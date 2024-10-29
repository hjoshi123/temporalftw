package datastore

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	"github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
