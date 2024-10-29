package helpers

import (
	"context"
	"encoding/json"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"math/rand/v2"
)

type BanksHelper struct {
	banksStore    datastoreIface.BankDataStore
	accountsStore datastoreIface.AccountStore
}

func NewBanksHelper() *BanksHelper {
	banksHelper := new(BanksHelper)
	banksHelper.banksStore = datastore.NewBankDataStore()
	banksHelper.accountsStore = datastore.NewAccountsStore()
	return banksHelper
}

func (bh *BanksHelper) CreateBank(ctx context.Context, bankInput *CreateBankInput) (*CreateBankOutput, error) {
	newBank := new(models.Bank)
	newBank.Name = bankInput.Name
	newBank.Address = bankInput.Address

	enc, err := json.Marshal(bankInput)
	if err != nil {
		return nil, err
	}

	newBank.Info = enc

	err = bh.banksStore.SaveBank(ctx, newBank)
	if err != nil {
		return nil, err
	}

	output := new(CreateBankOutput)
	output.ID = newBank.ID

	bankDefaultAccount := new(models.Account)
	bankDefaultAccount.BankID = newBank.ID
	bankDefaultAccount.AccountType = constants.BankDefaultAccount
	bankDefaultAccount.AccountNumber = rand.Int64()

	err = bh.accountsStore.SaveAccount(ctx, bankDefaultAccount)
	if err != nil {
		return nil, err
	}

	output.Msg = "Bank created successfully"

	return output, nil
}
