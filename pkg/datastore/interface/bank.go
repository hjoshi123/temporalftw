package datastore

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
)

type BankDataStore interface {
	SaveBank(ctx context.Context, bank *models.Bank) error
	GetBank(ctx context.Context, id int) (*models.Bank, error)
	GetAccountByNumberAndBankID(ctx context.Context, accountNumber int64, bankID int) (*models.Account, error)
	Withdraw(ctx context.Context, accountNumber int64, bankID int, amount float64) error
}
