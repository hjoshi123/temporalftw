package helpers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/internal/temporal"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	wfModels "github.com/hjoshi123/temporal-loan-app/pkg/models/workflow"
	"github.com/hjoshi123/temporal-loan-app/pkg/workflows"
	"go.temporal.io/sdk/client"
	temporal2 "go.temporal.io/sdk/temporal"
	"go.uber.org/zap"
)

type TXHelpers struct {
	accountStore datastoreIface.AccountStore
}

func NewTXHelpers() *TXHelpers {
	txHelpers := new(TXHelpers)
	txHelpers.accountStore = datastore.NewAccountsStore()
	return txHelpers
}

func (txh *TXHelpers) HandleStartTransaction(ctx context.Context, txInput *TransactionsInput) (*TransactionsOutput, error) {
	var bankID int

	switch txInput.TxType {
	case constants.TransactionTypeCredit:
		if txInput.ToBankID == 0 || txInput.ToAccountNumber == 0 {
			return nil, constants.ErrInvalidInputs
		}
		bankID = txInput.ToBankID
	case constants.TransactionTypeDebit:
		if txInput.FromBankID == 0 || txInput.FromAccountNumber == 0 {
			return nil, constants.ErrInvalidInputs
		}
		bankID = txInput.FromBankID
	}

	bankDefaultAccount, err := txh.accountStore.GetAccountByBankIDAndAccountType(ctx, bankID, constants.BankDefaultAccount)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank default account", "error", zap.Error(err))
		return nil, err
	}

	if txInput.FromBankID == 0 {
		txInput.FromBankID = bankDefaultAccount.BankID
	}

	if txInput.FromAccountNumber == 0 {
		txInput.FromAccountNumber = bankDefaultAccount.AccountNumber
	}

	if txInput.ToBankID == 0 {
		txInput.ToBankID = bankDefaultAccount.BankID
	}

	if txInput.ToAccountNumber == 0 {
		txInput.ToAccountNumber = bankDefaultAccount.AccountNumber
	}

	temporalClient, err := temporal.GetTemporalClient(ctx)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get temporal client", "error", zap.Error(err))
		return nil, err
	}

	defer temporalClient.Close()

	txUUID := uuid.New().String()

	startTxArgs := new(wfModels.TransactionWorkflowInput)
	startTxArgs.Amount = txInput.Amount
	startTxArgs.FromBankID = txInput.FromBankID
	startTxArgs.FromAccountNumber = txInput.FromAccountNumber
	startTxArgs.ToBankID = txInput.ToBankID
	startTxArgs.ToAccountNumber = txInput.ToAccountNumber
	startTxArgs.TxUUID = txUUID
	startTxArgs.TxType = txInput.TxType
	startTxArgs.Memo = txInput.Memo
	startTxArgs.TxStatus = constants.TransactionStatusPending

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        constants.CreateWorkflowID[string](constants.TransactionWorkflowName, txUUID),
		TaskQueue: constants.TransactionQueue,
		TypedSearchAttributes: temporal2.NewSearchAttributes(
			temporal2.NewSearchAttributeKeyString(constants.TransactionWorkflowSearchAttribute).ValueSet(txUUID),
		),
	}, workflows.TransactionWorkflow, startTxArgs)

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to start transaction workflow", "error", zap.Error(err))
		return nil, err
	}

	logging.FromContext(ctx).Infow("started transaction workflow", "workflowID", workflowRun.GetID(), "runID", workflowRun.GetRunID())

	output := new(TransactionsOutput)
	for {
		resp, err := temporalClient.QueryWorkflow(ctx, constants.CreateWorkflowID[string](constants.TransactionWorkflowName, txUUID), workflowRun.GetRunID(),
			constants.TransactionWorkflowStateIdentifier)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to query workflow", "error", zap.Error(err))
			return nil, err
		}

		activityOutput := new(wfModels.TransactionWorkflowOutput)
		err = resp.Get(&activityOutput)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to get query response", "error", zap.Error(err))
			return nil, err
		}

		if activityOutput.TransactionStatus == constants.TransactionStatusPending {
			output.TxUUID = activityOutput.TransactionID
			output.Msg = fmt.Sprintf("Tx with the ID %s is %s. Please approve it", activityOutput.TransactionID, constants.TransactionStatusPending)
			break
		} else if activityOutput.TransactionStatus == constants.TransactionStatusFailed {
			output.TxUUID = activityOutput.TransactionID
			output.Msg = fmt.Sprintf("Tx with the ID %s is rejected", activityOutput.TransactionID)
			break
		}
	}

	return output, nil
}

func (txh *TXHelpers) ApproveTransaction(ctx context.Context, approveTx *ApproveTransactionInput) error {
	temporalClient, err := temporal.GetTemporalClient(ctx)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get temporal client", "error", zap.Error(err))

	}

	defer temporalClient.Close()

	wflowID := constants.CreateWorkflowID[string](constants.TransactionWorkflowName, approveTx.TxUUID)
	err = temporalClient.SignalWorkflow(ctx, wflowID, "", constants.TransactionSignalApprove.String(), &wfModels.SignalApprove{
		TxApproved: true,
		TxUUID:     approveTx.TxUUID,
	})

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to signal", zap.Error(err))
		return err
	}

	return nil
}

func (txh *TXHelpers) RejectTransaction(ctx context.Context, rejectTx *RejectTransactionInput) error {
	temporalClient, err := temporal.GetTemporalClient(ctx)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get temporal client", "error", zap.Error(err))
	}

	defer temporalClient.Close()

	wflowID := constants.CreateWorkflowID[string](constants.TransactionWorkflowName, rejectTx.TxUUID)
	err = temporalClient.SignalWorkflow(ctx, wflowID, "", constants.TransactionSignalReject.String(), &wfModels.SignalReject{
		TxRejectedReason: rejectTx.RejectedReason,
		TxUUID:           rejectTx.TxUUID,
	})

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to signal", zap.Error(err))
		return err
	}

	return nil
}
