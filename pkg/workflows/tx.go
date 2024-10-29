package workflows

import (
	temporal2 "github.com/hjoshi123/temporal-loan-app/internal/temporal"
	"github.com/hjoshi123/temporal-loan-app/pkg/activities"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	wfModels "github.com/hjoshi123/temporal-loan-app/pkg/models/workflow"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

func TransactionWorkflow(ctx workflow.Context, txInput *wfModels.TransactionWorkflowInput) (*wfModels.TransactionWorkflowOutput, error) {
	result := new(wfModels.TransactionWorkflowOutput)
	commonActivities := activities.NewCommonActivities()

	ctx = temporal2.SetCommonActivityOptions(ctx)

	err := workflow.SetQueryHandler(ctx, constants.TransactionWorkflowStateIdentifier, func() (*wfModels.TransactionWorkflowOutput, error) {
		return result, nil
	})
	if err != nil {
		workflow.GetLogger(ctx).Error("failed to set query handler", zap.Error(err))
		result.Errors = append(result.Errors, err.Error())
		return result, err
	}

	detectFakeOutput := new(wfModels.FakeDetectionOutput)
	if err := workflow.ExecuteLocalActivity(ctx, commonActivities.DetectFakeOrInvalidData, txInput).
		Get(ctx, &detectFakeOutput); detectFakeOutput.IsFake {
		workflow.GetLogger(ctx).Error("failed to detect fake or invalid data", zap.Error(err))
		result.Errors = append(result.Errors, detectFakeOutput.FakeError.Error())
		return result, detectFakeOutput.FakeError
	}

	// Save transaction with pending state
	txActivities := activities.NewTransactionActivities()
	saveTxOutput := new(wfModels.SaveTransactionActivityOutput)
	if _ = workflow.ExecuteActivity(ctx, txActivities.SaveTransactionActivity, txInput).Get(ctx, &saveTxOutput); saveTxOutput.AbortTx {
		workflow.GetLogger(ctx).Error("failed to save transaction", zap.Error(saveTxOutput.AbortTxReason))
		result.Errors = append(result.Errors, saveTxOutput.AbortTxReason.Error())

		if err := workflow.ExecuteChildWorkflow(ctx, TransactionRejectWorkflow, wfModels.TransactionRejectWorkflowInput{
			TransactionID: saveTxOutput.TxUUID,
			Reason:        saveTxOutput.AbortTxReason.Error(),
		}).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("failed to execute rejection workflow", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
		}

		return result, saveTxOutput.AbortTxReason
	}

	result.TransactionStatus = saveTxOutput.TxStatus
	result.TransactionID = saveTxOutput.TxUUID
	result.FromAccountID = saveTxOutput.Tx.FromAccountNumber
	result.ToAccountID = saveTxOutput.Tx.ToAccountNumber

	rejectSignalChannel := workflow.GetSignalChannel(ctx, constants.TransactionSignalReject.String())
	approveSignalChannel := workflow.GetSignalChannel(ctx, constants.TransactionSignalApprove.String())

	// Once the transaction is saved with pending state now listen for events like approval or rejection of the transaction
	selector := workflow.NewSelector(ctx)

	selector.AddReceive(rejectSignalChannel, func(c workflow.ReceiveChannel, more bool) {
		var signal interface{}
		c.Receive(ctx, &signal)

		rejectData := new(wfModels.SignalReject)
		if err := mapstructure.Decode(signal, &rejectData); err != nil {
			workflow.GetLogger(ctx).Error("failed to decode signal", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
			return
		}

		if err := workflow.ExecuteChildWorkflow(temporal2.GetChildWorkflowOptions(ctx, saveTxOutput.TxUUID, constants.TransactionRejectWorkflowName),
			TransactionRejectWorkflow, &wfModels.TransactionRejectWorkflowInput{
				TransactionWorkflowInput: txInput,
				TransactionID:            saveTxOutput.TxUUID,
				Reason:                   rejectData.TxRejectedReason,
			}).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("failed to execute rejection workflow", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
		}

		result.TransactionStatus = constants.TransactionStatusFailed
	})

	selector.AddReceive(approveSignalChannel, func(c workflow.ReceiveChannel, more bool) {
		var signal interface{}
		c.Receive(ctx, &signal)

		approveData := new(wfModels.SignalApprove)
		if err := mapstructure.Decode(signal, &approveData); err != nil {
			workflow.GetLogger(ctx).Error("failed to decode signal", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
			return
		}

		if approveData.TxApproved {
			approveTxOutput := new(wfModels.TransactionApproveWorkflowOutput)
			if err := workflow.ExecuteChildWorkflow(temporal2.GetChildWorkflowOptions(ctx, saveTxOutput.TxUUID, constants.TransactionApproveWorkflowName),
				TransactionApproveWorkflow, &wfModels.TransactionApproveWorkflowInput{
					TransactionWorkflowInput: txInput,
					TransactionID:            approveData.TxUUID,
				}).Get(ctx, &approveTxOutput); err != nil {
				workflow.GetLogger(ctx).Error("failed to execute approval workflow", zap.Error(err))
				result.Errors = append(result.Errors, err.Error())
			}

			result.TransactionStatus = approveTxOutput.TxStatus
		}
	})

	selector.Select(ctx)

	return result, nil
}

func TransactionRejectWorkflow(ctx workflow.Context, txRejectInput *wfModels.TransactionRejectWorkflowInput) error {
	txRejectInput.TransactionWorkflowInput.TxStatus = constants.TransactionStatusFailed
	ctx = temporal2.SetCommonActivityOptions(ctx)

	txActivities := activities.NewTransactionActivities()
	saveTxOutput := new(wfModels.SaveTransactionActivityOutput)
	if err := workflow.ExecuteActivity(ctx, txActivities.SaveTransactionActivity, txRejectInput.TransactionWorkflowInput).
		Get(ctx, &saveTxOutput); err != nil {
		workflow.GetLogger(ctx).Error("failed to save transaction", zap.Error(saveTxOutput.AbortTxReason))

		return temporal.NewApplicationError(saveTxOutput.AbortTxReason.Error(), "")
	}

	return nil
}

func TransactionApproveWorkflow(ctx workflow.Context, txApproveInput *wfModels.TransactionApproveWorkflowInput) (*wfModels.TransactionApproveWorkflowOutput, error) {
	result := new(wfModels.TransactionApproveWorkflowOutput)
	txApproveInput.TransactionWorkflowInput.TxStatus = constants.TransactionStatusSuccess
	ctx = temporal2.SetCommonActivityOptions(ctx)

	txActivities := activities.NewTransactionActivities()
	saveTxOutput := new(wfModels.SaveTransactionActivityOutput)
	if err := workflow.ExecuteActivity(ctx, txActivities.SaveTransactionActivity, txApproveInput.TransactionWorkflowInput).
		Get(ctx, &saveTxOutput); err != nil {
		workflow.GetLogger(ctx).Error("failed to save transaction", zap.Error(saveTxOutput.AbortTxReason))

		return nil, temporal.NewApplicationError(saveTxOutput.AbortTxReason.Error(), "")
	}

	applyTxOutput := new(wfModels.ApplyTransactionActivityOutput)
	if err := workflow.ExecuteActivity(ctx, txActivities.ApplyTransactionActivity, &wfModels.ApplyTransactionActivityInput{
		TypeOfTransaction: txApproveInput.TxType,
		Tx:                saveTxOutput.Tx,
	}).Get(ctx, &applyTxOutput); err != nil {
		workflow.GetLogger(ctx).Error("failed to apply transaction", zap.Error(err))
		return nil, temporal.NewApplicationError(err.Error(), "")
	}

	workflow.GetLogger(ctx).Info("transaction applied successfully", zap.Any("transaction", applyTxOutput.ApplyStatus),
		zap.Any("transaction_id", applyTxOutput.TransactionID))

	result.TransactionID = applyTxOutput.TransactionID
	result.TxStatus = saveTxOutput.TxStatus

	return result, nil
}
