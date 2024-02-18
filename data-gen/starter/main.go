package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	datagen "github.com/nucleuscloud/temporal-feb24-meetup/data-gen"
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
		ID:        "datagen_" + uuid.NewString(),
		TaskQueue: "datagen",
	}

	ctx := context.Background()

	we, err := c.ExecuteWorkflow(ctx, workflowOptions, datagen.GenerateData)
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
