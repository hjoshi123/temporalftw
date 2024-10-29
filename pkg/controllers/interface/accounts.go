package controllers

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/handler"
)

type AccountsController interface {
	CreateAccount(ctx context.Context, input handler.Input) (handler.Output, error)
}
