package workflow

import (
	"awesomeProject2/activity"
	"awesomeProject2/model"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

const GalerryQueueName = "GALLERY_TASK_QUEUE"

func GalleryWorkFlow(ctx workflow.Context, user model.User, post model.PostItem) error {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    0,
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	}

	//apply the options for workflow
	ctx = workflow.WithActivityOptions(ctx, options)

	// Execute login activity
	err := workflow.ExecuteActivity(ctx, activity.Login, user).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Execute post image activity
	err = workflow.ExecuteActivity(ctx, activity.PostImage, post).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Execute share post activity
	err = workflow.ExecuteActivity(ctx, activity.SharePost, post).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
