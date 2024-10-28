package controllers

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/server"
)

// TransactionsController is the interface for the transactions controller. These methods are to be implemented
// by each version of the transactions' controller. If a method is not implemented in the previous version return an error to the user.
type TransactionsController interface {
	StartTransaction(ctx context.Context, input server.Input) (server.Output, error)
	ApproveTransaction(ctx context.Context, input server.Input) (server.Output, error)
	RejectTransaction(ctx context.Context, input server.Input) (server.Output, error)
	GetTransaction(ctx context.Context, input server.Input) (server.Output, error)
	GetTransactions(ctx context.Context, input server.Input) (server.Output, error)
}
