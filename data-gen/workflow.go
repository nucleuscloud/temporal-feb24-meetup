package datagen

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GenerateData(ctx workflow.Context) (any, error) {
	logger := workflow.GetLogger(ctx)
	_ = logger

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	var a *Activities

	// Executes first activity to retrieve data generation config
	// This will contain info like the database connection, as well as how many items to generate and each method to run per column
	var genConfig string
	err := workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.GetGenerateConfig,
	).Get(ctx, &genConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to execute GetGenerateConfig: %w", err)
	}

	// Executes the second activity that takes this configuration and simply runs it against the destination data source
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
