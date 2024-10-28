package temporal

import (
	"context"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"

	"logur.dev/logur"
)

func GetTemporalClient(ctx context.Context) (client.Client, error) {
	logger := logur.LoggerToKV(zapadapter.New(logging.FromContext(ctx).Desugar()))

	clientOptions := client.Options{
		Logger: logger,
		DataConverter: converter.NewCompositeDataConverter(
			converter.NewNilPayloadConverter(),
			converter.NewByteSlicePayloadConverter(),
			converter.NewProtoJSONPayloadConverter(),
			converter.NewProtoPayloadConverter(),
		),
	}

	sdkClient, err := client.Dial(clientOptions)

	if err != nil {
		logging.FromContext(ctx).Errorw("failed to create temporal client", "error", zap.Error(err))
		return nil, err
	}

	return sdkClient, nil
}
