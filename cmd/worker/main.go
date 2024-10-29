package main

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/config"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/internal/temporal"
	"github.com/hjoshi123/temporal-loan-app/pkg/activities"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	"github.com/hjoshi123/temporal-loan-app/pkg/workflows"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logging.FromContext(ctx).Infow("starting worker init")
	logging.FromContext(ctx).Infow("config", zap.Any("spec", config.Spec))

	_ = database.Connect(ctx)

	temporalClient, err := temporal.GetTemporalClient(ctx)
	if err != nil {
		logging.FromContext(ctx).Fatalw("failed to get temporal client", "error", err)
	}

	defer temporalClient.Close()

	temporalWorker := worker.New(temporalClient, constants.TransactionQueue, worker.Options{})

	temporalWorker.RegisterWorkflow(workflows.TransactionWorkflow)
	temporalWorker.RegisterWorkflow(workflows.TransactionApproveWorkflow)
	temporalWorker.RegisterWorkflow(workflows.TransactionRejectWorkflow)

	txActivities := activities.NewTransactionActivities()
	temporalWorker.RegisterActivityWithOptions(txActivities, activity.RegisterOptions{})

	commonActivities := activities.NewCommonActivities()
	temporalWorker.RegisterActivityWithOptions(commonActivities, activity.RegisterOptions{})

	logging.FromContext(ctx).Infow("starting worker", "queue", constants.TransactionQueue)
	err = temporalWorker.Run(worker.InterruptCh())
	if err != nil {
		logging.FromContext(ctx).Fatalw("failed to start worker", "error", err)
	}
}
