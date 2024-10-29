package v1

import (
	"context"
	"encoding/json"
	"github.com/hjoshi123/temporal-loan-app/internal/handler"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	controllers "github.com/hjoshi123/temporal-loan-app/pkg/controllers/interface"
	helpers "github.com/hjoshi123/temporal-loan-app/pkg/helpers/business/transactions"
)

type v1TransactionsController struct {
	txHelper *helpers.TXHelpers
}

func NewV1TransactionsController() controllers.TransactionsController {
	v1TC := new(v1TransactionsController)
	v1TC.txHelper = helpers.NewTXHelpers()
	return v1TC
}

func (v1TC *v1TransactionsController) StartTransaction(ctx context.Context, input handler.Input) (handler.Output, error) {
	txParams := new(helpers.TransactionsInput)
	err := json.NewDecoder(input.R.Body).Decode(&txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return handler.Output{}, err
	}

	txOutput, err := v1TC.txHelper.HandleStartTransaction(ctx, txParams)
	if err != nil {
		return handler.Output{}, err
	}

	return handler.Output{
		Output: txOutput,
	}, nil
}

func (v1TC *v1TransactionsController) ApproveTransaction(ctx context.Context, input handler.Input) (handler.Output, error) {
	txParams := new(helpers.ApproveTransactionInput)
	err := json.NewDecoder(input.R.Body).Decode(&txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return handler.Output{}, err
	}

	err = v1TC.txHelper.ApproveTransaction(ctx, txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to approve transaction", "error", err)
		return handler.Output{}, err
	}

	return handler.Output{}, nil
}

func (v1TC *v1TransactionsController) RejectTransaction(ctx context.Context, input handler.Input) (handler.Output, error) {
	txParams := new(helpers.RejectTransactionInput)
	err := json.NewDecoder(input.R.Body).Decode(&txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return handler.Output{}, err
	}

	err = v1TC.txHelper.RejectTransaction(ctx, txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to reject transaction", "error", err)
		return handler.Output{}, err
	}

	return handler.Output{}, nil
}
