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
	bankStore     datastoreIface.BankDataStore
	accountsStore datastoreIface.AccountStore
}

func NewCommonActivities() *CommonActivities {
	commonActivities := new(CommonActivities)
	commonActivities.bankStore = datastore.NewBankDataStore()
	commonActivities.accountsStore = datastore.NewAccountsStore()
	return commonActivities
}

func (ca *CommonActivities) DetectFakeOrInvalidData(ctx context.Context, input *models.TransactionWorkflowInput) (*models.FakeDetectionOutput, error) {
	output := new(models.FakeDetectionOutput)
	// Get the bank details from the bank store
	_, err := ca.bankStore.GetBank(ctx, input.FromBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank", "error", zap.Error(err))
		output.IsFake = true
		output.FakeError = constants.ErrFakeOrInvalidBank
		return output, nil
	}

	_, err = ca.bankStore.GetBank(ctx, input.ToBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank", "error", zap.Error(err))
		output.IsFake = true
		output.FakeError = constants.ErrFakeOrInvalidBank
		return output, nil
	}

	// Get the account details from the bank store
	_, err = ca.accountsStore.GetAccountByNumberAndBankID(ctx, input.FromAccountNumber, input.FromBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", zap.Error(err))
		output.IsFake = true
		output.FakeError = constants.ErrFakeOrInvalidAccount
		return output, nil
	}

	_, err = ca.accountsStore.GetAccountByNumberAndBankID(ctx, input.ToAccountNumber, input.ToBankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", zap.Error(err))
		output.IsFake = true
		output.FakeError = constants.ErrFakeOrInvalidAccount
		return output, nil
	}

	output.IsFake = false
	return output, nil
}
