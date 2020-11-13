package main

import (
	"flag"
	"fmt"

	"github.com/stmcginnis/gofish"

	"github.com/dhellmann/redfish-event-experiment/config"
)

func main() {
	flag.Parse()

	appConfig, err := config.LoadFromFile("config.yaml")
	if err != nil {
		panic(err)
	}

	gofishConfig := gofish.ClientConfig{
		Endpoint:  appConfig.BMC.URL,
		Username:  appConfig.BMC.Username,
		Password:  appConfig.BMC.Password,
		Insecure:  true,
		BasicAuth: true,
	}

	c, err := gofish.Connect(gofishConfig)
	if err != nil {
		panic(err)
	}
	service := c.Service

	es, err := service.EventService()
	if err != nil {
		panic(err)
	}

	for _, sub := range flag.Args() {
		err = es.DeleteEventSubscription(sub)
		if err != nil {
			panic(err)
		}
		fmt.Printf("deleted %s\n", sub)
	}
}
