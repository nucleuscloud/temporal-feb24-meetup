package datagen

import (
	"context"
	"fmt"
	"os"

	_ "github.com/benthosdev/benthos/v4/public/components/pure"
	_ "github.com/benthosdev/benthos/v4/public/components/pure/extended"
	_ "github.com/benthosdev/benthos/v4/public/components/sql"

	"github.com/benthosdev/benthos/v4/public/service"
	"go.temporal.io/sdk/activity"
)

type Activities struct{}

func (a *Activities) GetGenerateConfig() (string, error) {
	bits, err := os.ReadFile("./benthos.yaml")
	if err != nil {
		return "", fmt.Errorf("unable to read configuration file: %w", err)
	}
	return string(bits), nil
}

func (a *Activities) SynchronizeTable(
	ctx context.Context,
	config string,
) error {
	logger := activity.GetLogger(ctx)
	stream, err := getStreamFromConfig(config)
	if err != nil {
		return err
	}

	err = stream.Run(ctx)
	if err != nil {
		return fmt.Errorf("unable to run stream for sync: %w", err)
	}
	logger.Info("SynchronizeTable activity completed successfully")
	return nil
}

func getStreamFromConfig(config string) (*service.Stream, error) {
	streambldr := service.NewStreamBuilder()
	err := streambldr.SetYAML(config)
	if err != nil {
		return nil, fmt.Errorf("unable to set yaml for sync: %w", err)
	}

	stream, err := streambldr.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to build stream for sync: %w", err)
	}
	return stream, nil
}
