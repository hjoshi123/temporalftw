package v1

import (
	"context"
	"encoding/json"
	"github.com/hjoshi123/temporal-loan-app/internal/handler"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	controllers "github.com/hjoshi123/temporal-loan-app/pkg/controllers/interface"
	helpers "github.com/hjoshi123/temporal-loan-app/pkg/helpers/business/accounts"
)

type v1AccountsController struct {
	accountsHelper *helpers.AccountHelpers
}

func NewV1AccountsController() controllers.AccountsController {
	v1AC := new(v1AccountsController)
	v1AC.accountsHelper = helpers.NewAccountHelpers()
	return v1AC
}

func (v1AC *v1AccountsController) CreateAccount(ctx context.Context, input handler.Input) (handler.Output, error) {
	accountInput := new(helpers.CreateAccountInput)
	err := json.NewDecoder(input.R.Body).Decode(&accountInput)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to decode request body", "error", err)
		return handler.Output{}, err
	}

	createAcOutput, err := v1AC.accountsHelper.CreateAccount(ctx, accountInput)
	if err != nil {
		return handler.Output{}, err
	}

	return handler.Output{
		Output: createAcOutput,
	}, nil
}
