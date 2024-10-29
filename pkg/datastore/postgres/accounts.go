package datastore

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type accountsStore struct{}

func NewAccountsStore() datastoreIface.AccountStore {
	return new(accountsStore)
}

func (ac *accountsStore) SaveAccount(ctx context.Context, account *models.Account) error {
	db := database.Connect(ctx)

	err := account.Insert(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to save account", "error", err)
		return err
	}

	return nil
}

func (ac *accountsStore) GetAccountByNumberAndBankID(ctx context.Context, accountNumber int64, bankID int) (*models.Account, error) {
	db := database.Connect(ctx)

	account, err := models.Accounts(models.AccountWhere.AccountNumber.EQ(accountNumber), models.AccountWhere.BankID.EQ(bankID)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return nil, err
	}

	return account, nil
}

func (ac *accountsStore) GetAccountByBankIDAndAccountType(ctx context.Context, bankID int, accountType string) (*models.Account, error) {
	db := database.Connect(ctx)

	account, err := models.Accounts(models.AccountWhere.BankID.EQ(bankID), models.AccountWhere.AccountType.EQ(accountType)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return nil, err
	}

	return account, nil
}

func (ac *accountsStore) Withdraw(ctx context.Context, accountNumber int64, bankID int, amount types.Decimal) error {
	db := database.Connect(ctx)

	account, err := ac.GetAccountByNumberAndBankID(ctx, accountNumber, bankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return err
	}

	account.Balance = types.NewDecimal(account.Balance.Sub(account.Balance.Big, amount.Big))

	_, err = account.Update(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to update account", "error", err)
		return err
	}

	return nil
}

func (ac *accountsStore) Deposit(ctx context.Context, accountNumber int64, bankID int, amount types.Decimal) error {
	db := database.Connect(ctx)

	account, err := ac.GetAccountByNumberAndBankID(ctx, accountNumber, bankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return err
	}

	account.Balance = types.NewDecimal(account.Balance.Add(account.Balance.Big, amount.Big))

	_, err = account.Update(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to update account", "error", err)
		return err
	}

	return nil
}
