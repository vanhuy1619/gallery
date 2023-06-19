package activity

import (
	"awesomeProject2/model"
	"context"
	"go.temporal.io/sdk/activity"
	"log"
)

type User model.User

type PostItem model.PostItem

// Login activity
// Call Adapter
func Login(ctx context.Context, user User) error {
	info := activity.GetInfo(ctx)
	workflowId := info.WorkflowExecution.ID
	runId := info.WorkflowExecution.RunID
	log.Printf("Start activity: UploadImgTrustingSocial, workflowId: %s, runId: %s", workflowId, runId)
	log.Printf("Performing login activity for user: %s\n", user.Username)
	// Perform login logic here
	//repositories.Login
	return nil
}

// PostImage activity
func PostImage(ctx context.Context, post PostItem) error {
	log.Printf("Performing post image activity")
	// Perform post image logic here
	return nil
}

// SharePost activity
func SharePost(ctx context.Context, post PostItem) error {
	log.Printf("Performing share post activity")
	// Perform share post logic here
	return nil
}
