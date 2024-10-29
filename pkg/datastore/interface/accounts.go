package datastore

import (
	"context"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type AccountStore interface {
	SaveAccount(ctx context.Context, account *models.Account) error
	GetAccountByNumberAndBankID(ctx context.Context, accountNumber int64, bankID int) (*models.Account, error)
	Withdraw(ctx context.Context, accountNumber int64, bankID int, amount types.Decimal) error
	Deposit(ctx context.Context, accountNumber int64, bankID int, amount types.Decimal) error
	GetAccountByBankIDAndAccountType(ctx context.Context, bankID int, accountType string) (*models.Account, error)
}
