package v1

import (
	"context"
	"encoding/json"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/internal/server"
	controllers "github.com/hjoshi123/temporal-loan-app/pkg/controllers/interface"
	helpers "github.com/hjoshi123/temporal-loan-app/pkg/helpers/transactions"
)

type v1TransactionsController struct {
	txHelper *helpers.TXHelpers
}

func NewV1TransactionsController() controllers.TransactionsController {
	v1TC := new(v1TransactionsController)
	v1TC.txHelper = helpers.NewTXHelpers()
	return v1TC
}

func (v1TC *v1TransactionsController) StartTransaction(ctx context.Context, input server.Input) (server.Output, error) {
	txParams := new(helpers.TransactionsInput)
	err := json.NewDecoder(input.R.Body).Decode(&txParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return server.Output{}, err
	}

}
