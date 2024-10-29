package activities

import (
	"context"
	"fmt"
	"github.com/ericlagergren/decimal"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	wfModels "github.com/hjoshi123/temporal-loan-app/pkg/models/workflow"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"go.temporal.io/sdk/temporal"
	"math"
)

type TransactionActivities struct {
	txStore      datastoreIface.TransactionStore
	bankStore    datastoreIface.BankDataStore
	accountStore datastoreIface.AccountStore
}

func NewTransactionActivities() *TransactionActivities {
	txActivities := new(TransactionActivities)
	txActivities.txStore = datastore.NewTransactionStore()
	txActivities.bankStore = datastore.NewBankDataStore()
	txActivities.accountStore = datastore.NewAccountsStore()
	return txActivities
}

func (txAc *TransactionActivities) SaveTransactionActivity(ctx context.Context, txInput *wfModels.TransactionWorkflowInput) (*wfModels.SaveTransactionActivityOutput, error) {
	output := new(wfModels.SaveTransactionActivityOutput)

	if txInput.Amount < 0 {
		output.AbortTx = true
		output.AbortTxReason = constants.ErrInvalidAmount
		return output, nil
	}

	decimalScaleAmount := decimal.New(int64(txInput.Amount*math.Pow10(6)), 6)
	decimalAmount := types.NewDecimal(decimalScaleAmount)

	txUUIDString := null.NewString(txInput.TxUUID, true)

	err := txAc.txStore.SaveTransaction(ctx, &models.Transaction{
		Amount:            decimalAmount,
		FromAccountNumber: txInput.FromAccountNumber,
		FromBankID:        txInput.FromBankID,
		ToAccountNumber:   txInput.ToAccountNumber,
		ToBankID:          txInput.ToBankID,
		TransactionID:     txUUIDString,
		Description:       txInput.Memo,
	}, txInput.TxType, txInput.TxStatus)
	if err != nil {
		output.AbortTx = true
		output.AbortTxReason = fmt.Errorf("failed to save transaction: %w", err)
		return output, nil
	}

	savedTx, err := txAc.txStore.GetTransactionByUUID(ctx, txInput.TxUUID)
	if err != nil {
		output.AbortTx = true
		output.AbortTxReason = fmt.Errorf("failed to get saved transaction: %w", err)
		return output, nil
	}

	output.TxUUID = savedTx.TransactionID.String
	output.TxStatus = savedTx.R.TransactionStatus.Name
	output.Tx = savedTx

	return output, nil
}

// ApplyTransactionActivity only supports one way transactions i.e. either from the bank account to the user account or vice versa.
// It does not support transactions between two bank accounts.
func (txAc *TransactionActivities) ApplyTransactionActivity(ctx context.Context, applyTx *wfModels.ApplyTransactionActivityInput) (*wfModels.ApplyTransactionActivityOutput, error) {
	output := new(wfModels.ApplyTransactionActivityOutput)

	switch applyTx.TypeOfTransaction {
	case constants.TransactionTypeCredit:
		logging.FromContext(ctx).Infow("credit transaction started", "transaction_id", applyTx.Tx.TransactionID)
		// Credit the amount to the account

		err := txAc.accountStore.Deposit(ctx, applyTx.Tx.ToAccountNumber, applyTx.Tx.ToBankID, applyTx.Tx.Amount)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to credit amount", "error", err)
			output.ApplyStatus = constants.TransactionStatusFailed
			output.TransactionID = applyTx.Tx.TransactionID.String
			output.TypeOfTransaction = applyTx.TypeOfTransaction
			return output, temporal.NewApplicationError("failed to credit amount", "")
		}

		output.ApplyStatus = constants.TransactionStatusSuccess
		output.TransactionID = applyTx.Tx.TransactionID.String
		output.TypeOfTransaction = applyTx.TypeOfTransaction
		logging.FromContext(ctx).Infow("credit transaction completed", "transaction_id", applyTx.Tx.TransactionID)

		return output, nil
	case constants.TransactionTypeDebit:
		logging.FromContext(ctx).Infow("debit transaction started", "transaction_id", applyTx.Tx.TransactionID)
		// Debit the amount from the account

		err := txAc.accountStore.Withdraw(ctx, applyTx.Tx.FromAccountNumber, applyTx.Tx.FromBankID, applyTx.Tx.Amount)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to credit amount", "error", err)
			output.ApplyStatus = constants.TransactionStatusFailed
			output.TransactionID = applyTx.Tx.TransactionID.String
			output.TypeOfTransaction = applyTx.TypeOfTransaction
			return output, temporal.NewApplicationError("failed to credit amount", "")
		}

		output.ApplyStatus = constants.TransactionStatusSuccess
		output.TransactionID = applyTx.Tx.TransactionID.String
		output.TypeOfTransaction = applyTx.TypeOfTransaction
		logging.FromContext(ctx).Infow("credit transaction completed", "transaction_id", applyTx.Tx.TransactionID)

		return output, nil
	}

	output.ApplyStatus = "invalid transaction type"
	output.TransactionID = applyTx.Tx.TransactionID.String
	output.TypeOfTransaction = applyTx.TypeOfTransaction
	return output, nil
}
