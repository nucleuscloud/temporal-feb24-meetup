package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	mldatagen "github.com/nucleuscloud/temporal-feb24-meetup/ml-data-gen"
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
		"mldatagen",
		worker.Options{},
	)

	w.RegisterWorkflow(mldatagen.TrainModel)
	w.RegisterWorkflow(mldatagen.SampleModel)
	w.RegisterActivity(&mldatagen.Activities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
