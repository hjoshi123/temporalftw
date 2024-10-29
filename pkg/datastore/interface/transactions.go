package datastore

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
)

type TransactionStore interface {
	SaveTransaction(ctx context.Context, transaction *models.Transaction, transactionType, transactionStatus string) error
	GetTransaction(ctx context.Context, id int) (*models.Transaction, error)
	GetTransactions(ctx context.Context, accountNumber int64, bankID int) ([]*models.Transaction, error)
	GetTransactionByUUID(ctx context.Context, uuid string) (*models.Transaction, error)
}
