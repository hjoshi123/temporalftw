package datastore

import (
	"context"
	"github.com/ericlagergren/decimal"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	"github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
	"math"
)

type postgresDB struct{}

func NewBankDataStore() datastoreIface.BankDataStore {
	return &postgresDB{}
}

func (p *postgresDB) SaveBank(ctx context.Context, bank *models.Bank) error {
	db := database.Connect(ctx)

	err := bank.Insert(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to save bank", "error", err)
		return err
	}

	return nil
}

func (p *postgresDB) GetBank(ctx context.Context, id int) (*models.Bank, error) {
	db := database.Connect(ctx)

	bank, err := models.Banks(models.BankWhere.ID.EQ(id)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get bank", "error", err)
		return nil, err
	}

	return bank, nil
}

func (p *postgresDB) GetAccountByNumberAndBankID(ctx context.Context, accountNumber int64, bankID int) (*models.Account, error) {
	db := database.Connect(ctx)

	account, err := models.Accounts(models.AccountWhere.AccountNumber.EQ(accountNumber), models.AccountWhere.BankID.EQ(bankID)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return nil, err
	}

	return account, nil
}

func (p *postgresDB) Withdraw(ctx context.Context, accountNumber int64, bankID int, amount float64) error {
	db := database.Connect(ctx)

	account, err := p.GetAccountByNumberAndBankID(ctx, accountNumber, bankID)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get account", "error", err)
		return err
	}

	decimalScaleAmount := decimal.New(int64(amount*math.Pow10(6)), 6)
	decimalAmount := types.NewDecimal(decimalScaleAmount)

	account.Balance = types.NewDecimal(account.Balance.Sub(account.Balance.Big, decimalAmount.Big))

	_, err = account.Update(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to update account", "error", err)
		return err
	}

	return nil
}
