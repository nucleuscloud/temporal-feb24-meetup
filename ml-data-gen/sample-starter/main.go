package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	mldatagen "github.com/nucleuscloud/temporal-feb24-meetup/ml-data-gen"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("nnable to create temporal client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:                 "mldatagen_sampler" + uuid.NewString(),
		TaskQueue:          "mldatagen",
		WorkflowRunTimeout: 30 * time.Second,
	}

	ctx := context.Background()

	we, err := c.ExecuteWorkflow(ctx, workflowOptions, mldatagen.SampleModel)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}
	log.Println("started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(ctx, &result)
	if err != nil {
		log.Fatalln("unable get workflow result", err)
	}
	log.Println("workflow result:", result)
}
