package workflows

import (
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

	if err := workflow.ExecuteLocalActivity(ctx, commonActivities.DetectFakeOrInvalidData, txInput).Get(ctx, nil); err != nil {
		workflow.GetLogger(ctx).Error("failed to detect fake or invalid data", zap.Error(err))
		result.Errors = append(result.Errors, err.Error())
		return result, err
	}

	// Save transaction with pending state
	txActivities := activities.NewTransactionActivities()
	saveTxOutput := new(wfModels.SaveTransactionActivityOutput)
	if _ = workflow.ExecuteActivity(ctx, txActivities.SaveTransactionActivity, txInput).Get(ctx, &saveTxOutput); saveTxOutput.AbortTx {
		workflow.GetLogger(ctx).Error("failed to save transaction", zap.Error(saveTxOutput.AbortTxReason))
		result.Errors = append(result.Errors, saveTxOutput.AbortTxReason.Error())

		if err := workflow.ExecuteChildWorkflow(ctx, TransactionRejectWorkflow, &wfModels.TransactionRejectWorkflowInput{
			TransactionID: saveTxOutput.TxUUID,
			Reason:        saveTxOutput.AbortTxReason.Error(),
		}).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("failed to execute rejection workflow", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
		}

		return result, saveTxOutput.AbortTxReason
	}

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

		if err := workflow.ExecuteChildWorkflow(ctx, TransactionRejectWorkflow, &wfModels.TransactionRejectWorkflowInput{
			TransactionID: saveTxOutput.TxUUID,
			Reason:        rejectData.TxRejectedReason,
		}).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("failed to execute rejection workflow", zap.Error(err))
			result.Errors = append(result.Errors, err.Error())
		}
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
			if err := workflow.ExecuteChildWorkflow(ctx, TransactionApproveWorkflow, &wfModels.TransactionApproveWorkflowInput{
				TransactionWorkflowInput: txInput,
				TransactionID:            saveTxOutput.TxUUID,
			}).Get(ctx, nil); err != nil {
				workflow.GetLogger(ctx).Error("failed to execute approval workflow", zap.Error(err))
				result.Errors = append(result.Errors, err.Error())
			}
		}
	})

	selector.Select(ctx)

	return result, nil
}

func TransactionRejectWorkflow(ctx workflow.Context, txRejectInput *wfModels.TransactionRejectWorkflowInput) {
}

func TransactionApproveWorkflow(ctx workflow.Context, txApproveInput *wfModels.TransactionApproveWorkflowInput) error {
	txApproveInput.TransactionWorkflowInput.TxType = constants.TransactionStatusSuccess

	txActivities := activities.NewTransactionActivities()
	saveTxOutput := new(activities.SaveTransactionActivityOutput)
	if err := workflow.ExecuteActivity(ctx, txActivities.SaveTransactionActivity, txApproveInput.TransactionWorkflowInput).
		Get(ctx, &saveTxOutput); err != nil {
		workflow.GetLogger(ctx).Error("failed to save transaction", zap.Error(saveTxOutput.AbortTxReason))

		return temporal.NewApplicationError(saveTxOutput.AbortTxReason.Error(), "")
	}

}
