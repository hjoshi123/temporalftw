package helpers

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"math/rand/v2"
)

type AccountHelpers struct {
	accountStore datastoreIface.AccountStore
}

func NewAccountHelpers() *AccountHelpers {
	acHelpers := new(AccountHelpers)
	acHelpers.accountStore = datastore.NewAccountsStore()
	return acHelpers
}

func (ac *AccountHelpers) CreateAccount(ctx context.Context, input *CreateAccountInput) (*CreateAccountOutput, error) {
	output := new(CreateAccountOutput)
	if input.BankID == 0 {
		return nil, constants.ErrInvalidInputs
	}

	account := new(models.Account)
	account.AccountNumber = rand.Int64()
	account.BankID = input.BankID
	account.AccountType = input.AccountType

	err := ac.accountStore.SaveAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	output.AccountNumber = account.AccountNumber
	output.BankID = account.BankID
	return output, nil
}
