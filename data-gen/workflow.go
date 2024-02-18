package datagen

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func GenerateData(ctx workflow.Context) (any, error) {
	logger := workflow.GetLogger(ctx)
	_ = logger

	ao := workflow.ActivityOptions{}

	var a *Activities

	var genConfig string
	err := workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.GetGenerateConfig,
	).Get(ctx, &genConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to execute GetGenerateConfig: %w", err)
	}

	err = workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.SynchronizeTable,
		genConfig,
	).Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to execute SynchronizeTable: %w", err)
	}

	logger.Info("GenerateData workflow completed successfully")
	return nil, nil
}
