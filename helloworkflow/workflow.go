package helloworkflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, name string) (string, error) { // Replace name string with a param struct
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}

	// Set maximum retry and also listen for specific errors

	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed", "Error", err)
	}

	return result, nil
}

func Activity(ctx context.Context, name string) (string, error) {
	return "Hello " + name, nil
}
