package v1

import (
	"context"
	"encoding/json"
	"github.com/hjoshi123/temporal-loan-app/internal/handler"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	controllers "github.com/hjoshi123/temporal-loan-app/pkg/controllers/interface"
	helpers "github.com/hjoshi123/temporal-loan-app/pkg/helpers/business/banks"
)

type v1BanksController struct {
	banksHelper *helpers.BanksHelper
}

func NewV1BanksController() controllers.BanksController {
	v1BanksController := new(v1BanksController)
	v1BanksController.banksHelper = helpers.NewBanksHelper()
	return v1BanksController
}

func (bc *v1BanksController) CreateBank(ctx context.Context, input handler.Input) (handler.Output, error) {
	createBankParams := new(helpers.CreateBankInput)
	err := json.NewDecoder(input.R.Body).Decode(&createBankParams)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return handler.Output{}, err
	}

	createBankOutput, err := bc.banksHelper.CreateBank(ctx, createBankParams)
	if err != nil {
		return handler.Output{}, err
	}

	return handler.Output{
		Output: createBankOutput,
	}, nil
}
