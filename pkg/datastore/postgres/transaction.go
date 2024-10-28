package datastore

import (
	"context"
	"fmt"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	"github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type transactionPostgres struct{}

func NewTransactionStore() datastoreIface.TransactionStore {
	return &transactionPostgres{}
}

func (t *transactionPostgres) SaveTransaction(ctx context.Context, transaction *models.Transaction, transactionType string) error {
	db := database.Connect(ctx)

	txType, err := models.TransactionTypes(
		qm.Where(fmt.Sprintf("%s = %s", models.TransactionTypeColumns.Name, transactionType)),
	).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get transaction type", "error", err)
		return err
	}

	status, err := models.TransactionStatuses(
		qm.Where(fmt.Sprintf("%s = %s", models.TransactionStatusColumns.Name, constants.TransactionStatusPending)),
	).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get transaction status", "error", err)
		return err
	}

	transaction.TransactionStatusID = status.ID
	transaction.TransactionTypeID = txType.ID
	err = transaction.Insert(ctx, db, boil.Infer())
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to save transaction", "error", err)
		return err
	}

	return nil
}

func (t *transactionPostgres) GetTransaction(ctx context.Context, id int) (*models.Transaction, error) {
	db := database.Connect(ctx)

	tx, err := models.Transactions(models.TransactionWhere.ID.EQ(id)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get transaction", "error", err)
		return nil, err
	}

	return tx, nil
}

func (t *transactionPostgres) GetTransactionByUUID(ctx context.Context, uuid string) (*models.Transaction, error) {
	db := database.Connect(ctx)

	tx, err := models.Transactions(models.TransactionWhere.TransactionID.EQ(null.NewString(uuid, true)),
		qm.Load(models.TransactionRels.TransactionStatus)).One(ctx, db)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get transaction by uuid", "error", err)
		return nil, err
	}

	return tx, nil
}

func (t *transactionPostgres) GetTransactions(ctx context.Context, fromAccountNumber int64, bankID int) ([]*models.Transaction, error) {
	db := database.Connect(ctx)

	transactions, err := models.Transactions(models.TransactionWhere.FromAccountNumber.EQ(fromAccountNumber),
		models.TransactionWhere.FromBankID.EQ(bankID)).
		All(ctx, db)

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get transactions", "error", err)
		return nil, err
	}

	return transactions, nil
}
