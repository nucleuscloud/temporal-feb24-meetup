package mldatagen

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

const (
	// task queue specific to the python workers that train/sample ML models
	mlTaskQueue = "ml"
)

func TrainModel(ctx workflow.Context) (any, error) {
	logger := workflow.GetLogger(ctx)
	_ = logger

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	var a *Activities

	var getTrainInputResponse *GetTrainInputResponse
	err := workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.GetTrainInput,
	).Get(ctx, &getTrainInputResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to execute GetTrainInput activity: %w", err)
	}

	err = workflow.ExecuteActivity(
		workflow.WithTaskQueue(workflow.WithActivityOptions(ctx, ao), mlTaskQueue),
		"train-table-model",
		getTrainInputResponse.TrainModelRequest,
	).Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to execute train-table-model activity: %w", err)
	}

	logger.Info("TrainModel workflow completed successfully")
	return nil, nil
}

func SampleModel(ctx workflow.Context) (any, error) {
	logger := workflow.GetLogger(ctx)
	_ = logger

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	var a *Activities

	var getSampleInputResponse *GetSampleInputResponse
	err := workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.GetSampleInput,
	).Get(ctx, &getSampleInputResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to execute GetSampleInput activity: %w", err)
	}

	// upsert table schema
	err = workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, ao),
		a.UpsertTableSchema,
		&UpsertTableSchemaRequest{
			Dsn:    getSampleInputResponse.Dsn,
			Schema: getSampleInputResponse.Schema,
			Table:  getSampleInputResponse.Table,
		},
	).Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to execute UpsertTableSchema activity: %w", err)
	}

	err = workflow.ExecuteActivity(
		workflow.WithTaskQueue(workflow.WithActivityOptions(ctx, ao), mlTaskQueue),
		"sample-table-model",
		getSampleInputResponse.SampleModelRequest,
	).Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to execute sample-table-model activity: %w", err)
	}

	logger.Info("SampleModel workflow completed successfully")
	return nil, nil
}
