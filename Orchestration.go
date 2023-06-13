package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"time"
//
//	"github.com/gocql/gocql"
//	"github.com/segmentio/kafka-go"
//	"go.temporal.io/sdk/client"
//	"go.temporal.io/sdk/worker"
//	"go.temporal.io/sdk/workflow"
//)
//
//// Constants
//const (
//	TaskQueueName  = "upload-task-queue"
//	CassandraHosts = "127.0.0.1"
//	KafkaBrokers   = "localhost:9092"
//	KafkaTopic     = "upload-notifications"
//)
//
//// UploadWorkflow implements the upload workflow
//func UploadWorkflow(ctx workflow.Context, userID string) error {
//	// Execute the activities in parallel
//	errGroup := workflow.Go(ctx, func(ctx workflow.Context) error {
//		future1 := workflow.ExecuteActivity(ctx, UploadIDCardImageActivity, userID)
//		future2 := workflow.ExecuteActivity(ctx, UploadSelfieImageActivity, userID)
//		future3 := workflow.ExecuteActivity(ctx, UploadVideoActivity, userID)
//
//		err1 := future1.Get(ctx, nil)
//		err2 := future2.Get(ctx, nil)
//		err3 := future3.Get(ctx, nil)
//
//		if err1 != nil {
//			return err1
//		}
//		if err2 != nil {
//			return err2
//		}
//		if err3 != nil {
//			return err3
//		}
//
//		return nil
//	})
//
//	// Wait for all activities to complete
//	err := errGroup.Wait()
//	if err != nil {
//		return err
//	}
//
//	// Workflow completed successfully
//	return nil
//}
//
//// UploadIDCardImageActivity implements the activity to upload ID card image
//func UploadIDCardImageActivity(ctx context.Context, userID string) error {
//	// Simulate the upload process
//	time.Sleep(time.Second)
//	fmt.Println("Uploaded ID card image for user:", userID)
//
//	return nil
//}
//
//// UploadSelfieImageActivity implements the activity to upload selfie image
//func UploadSelfieImageActivity(ctx context.Context, userID string) error {
//	// Simulate the upload process
//	time.Sleep(time.Second)
//	fmt.Println("Uploaded selfie image for user:", userID)
//
//	return nil
//}
//
//// UploadVideoActivity implements the activity to upload video
//func UploadVideoActivity(ctx context.Context, userID string) error {
//	// Simulate the upload process
//	time.Sleep(time.Second)
//	fmt.Println("Uploaded video for user:", userID)
//
//	return nil
//}
//
//func main() {
//	// Create the Temporal client
//	c, err := client.NewClient(client.Options{})
//	if err != nil {
//		log.Fatal("Failed to create Temporal client:", err)
//	}
//
//	// Create the Temporal worker
//	w := worker.New(c, TaskQueueName, worker.Options{})
//
//	// Register the workflow and activity types with the worker
//	w.RegisterWorkflow(UploadWorkflow)
//	w.RegisterActivity(UploadIDCardImageActivity)
//	w.RegisterActivity(UploadSelfieImageActivity)
//	w.RegisterActivity(UploadVideoActivity)
//
//	// Start the worker
//	err = w.Run(worker.InterruptCh())
//	if err != nil {
//		log.Fatal("Failed to start Temporal worker:", err)
//	}
//
//	// Connect to Cassandra
//	cluster := gocql.NewCluster(CassandraHosts)
//	cluster.Keyspace = "uploads"
//	session, err := cluster.CreateSession()
//	if err != nil {
//		log.Fatal("Failed to connect to Cassandra:", err)
//	}
//	defer session.Close()
//
//	// Create the Kafka writer
//	writer := kafka.NewWriter(kafka.WriterConfig{
//		Brokers: []string{KafkaBrokers},
//		Topic:   KafkaTopic,
//	})
//
//	// Start a new workflow execution
//	exec, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
//		TaskQueue:           TaskQueueName,
//		WorkflowID:          "upload-workflow-" + time.Now().Format("20060102150405"),
//		WorkflowRunTimeout:  time.Minute * 5,
//		WorkflowTaskTimeout: time.Minute,
//	}, UploadWorkflow, "user123")
//
//	if err != nil {
//		log.Fatal("Failed to start workflow:", err)
//	}
//
//	// Wait for the workflow execution to complete
//	err = exec.Get(context.Background(), nil)
//	if err != nil {
//		log.Fatal("Failed to get workflow result:", err)
//	}
//
//	// Send a notification to Kafka
//	err = writer.WriteMessages(context.Background(), kafka.Message{
//		Key:   []byte("upload-complete"),
//		Value: []byte("Upload workflow completed"),
//	})
//	if err != nil {
//		log.Fatal("Failed to send Kafka message:", err)
//	}
//
//	// Close the Kafka writer
//	err = writer.Close()
//	if err != nil {
//		log.Fatal("Failed to close Kafka writer:", err)
//	}
//
//	fmt.Println("Upload workflow completed")
//}
