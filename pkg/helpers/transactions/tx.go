package helpers

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/internal/temporal"
	"go.uber.org/zap"
)

type TXHelpers struct{}

func NewTXHelpers() *TXHelpers {
	txHelpers := new(TXHelpers)
	return txHelpers
}

func (txh *TXHelpers) HandleStartTransaction(ctx context.Context, txInput *TransactionsInput) error {
	temporalClient, err := temporal.GetTemporalClient(ctx)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to get temporal client", "error", zap.Error(err))
		return err
	}

	defer temporalClient.Close()
}
