package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"time"
//
//	"go.temporal.io/sdk/client"
//	"go.temporal.io/sdk/worker"
//	"go.temporal.io/sdk/workflow"
//)
//
//// Workflow interface
//type TransferWorkflow interface {
//	TransferFunds(ctx workflow.Context, sourceAccountID string, destinationAccountID string, amount float64) error
//}
//
//// Workflow implementation
//func TransferFunds(ctx workflow.Context, sourceAccountID string, destinationAccountID string, amount float64) error {
//	// Simulate transferring funds from source account to destination account
//	fmt.Printf("Transferring %.2f from account %s to account %s\n", amount, sourceAccountID, destinationAccountID)
//
//	// TODO: Implement the logic to transfer funds in your banking system
//
//	// Simulating a delay to simulate the transfer process
//	time.Sleep(5 * time.Second)
//
//	// Check if the transfer was successful
//	successfulTransfer := true
//
//	if successfulTransfer {
//		fmt.Println("Transfer successful")
//		return nil
//	}
//
//	// If the transfer failed, return an error
//	return fmt.Errorf("transfer failed")
//}
//
//func main() {
//	// Create a new Temporal client
//	c, err := client.NewClient(client.Options{})
//	if err != nil {
//		log.Fatal("Unable to create Temporal client", err)
//	}
//
//	// Create a new Temporal worker
//	w := worker.New(c, "banking-tasks", worker.Options{})
//
//	// Register the TransferFunds workflow implementation with the worker
//	w.RegisterWorkflow(TransferFunds)
//
//	// Start the worker
//	err = w.Run(worker.InterruptCh())
//	if err != nil {
//		log.Fatal("Unable to start Temporal worker", err)
//	}
//
//	// Create a new workflow options
//	wo := client.StartWorkflowOptions{
//		ID:        "transaction-123",
//		TaskQueue: "banking-tasks",
//	}
//
//	// Start a new transfer workflow
//	_, err = c.ExecuteWorkflow(context.Background(), wo, TransferFunds, "source-account-123", "destination-account-456", 100.0)
//	if err != nil {
//		log.Fatal("Unable to start transfer workflow", err)
//	}
//
//	// Wait for the workflow execution to complete
//	fmt.Println("Waiting for workflow execution to complete...")
//}
