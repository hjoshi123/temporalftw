package activities

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/workflow"
	"go.uber.org/zap"
)

type CommonActivities struct {
	bankStore datastoreIface.BankDataStore
}

func NewCommonActivities() *CommonActivities {
	commonActivities := new(CommonActivities)
	commonActivities.bankStore = datastore.NewBankDataStore()
	return commonActivities
}

func (ca *CommonActivities) DetectFakeOrInvalidData(ctx context.Context, input *models.TransactionWorkflowInput) error {
	// Get the bank details from the bank store
	_, err := ca.bankStore.GetBank(ctx, input.FromBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank", "error", zap.Error(err))
		return constants.ErrFakeOrInvalidBank
	}

	_, err = ca.bankStore.GetBank(ctx, input.ToBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank", "error", zap.Error(err))
		return constants.ErrFakeOrInvalidBank
	}

	// Get the account details from the bank store
	_, err = ca.bankStore.GetAccountByNumberAndBankID(ctx, input.FromAccountNumber, input.FromBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", zap.Error(err))
		return constants.ErrFakeOrInvalidAccount
	}

	_, err = ca.bankStore.GetAccountByNumberAndBankID(ctx, input.ToAccountNumber, input.ToBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", zap.Error(err))
		return constants.ErrFakeOrInvalidAccount
	}

	return nil
}
