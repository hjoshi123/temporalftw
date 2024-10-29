package temporal

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/config"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
	"time"

	"logur.dev/logur"
)

func GetTemporalClient(ctx context.Context) (client.Client, error) {
	logger := logur.LoggerToKV(zapadapter.New(logging.FromContext(ctx).Desugar()))

	clientOptions := client.Options{
		HostPort: config.Spec.TemporalHostPort,
		Logger:   logger,
	}

	sdkClient, err := client.Dial(clientOptions)

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to create temporal client", "error", zap.Error(err))
		return nil, err
	}

	return sdkClient, nil
}

func SetCommonActivityOptions(ctx workflow.Context) workflow.Context {
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: time.Second * 10,
	})

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	})
	return ctx
}
