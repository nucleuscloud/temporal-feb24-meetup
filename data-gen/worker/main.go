package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	datagen "github.com/nucleuscloud/temporal-feb24-meetup/data-gen"
)

func main() {
	temporalclient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer temporalclient.Close()

	w := worker.New(
		temporalclient,
		"datagen",
		worker.Options{},
	)

	w.RegisterWorkflow(datagen.GenerateData)
	activities := &datagen.Activities{}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
