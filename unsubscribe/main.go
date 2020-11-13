package main

import (
	"flag"
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	var subscription string
	config := gofish.ClientConfig{
		Insecure:  true,
		BasicAuth: true,
	}
	flag.StringVar(&config.Endpoint, "endpoint", "", "The URL of the BMC")
	flag.StringVar(&config.Username, "user", "", "The username for authentication")
	flag.StringVar(&config.Password, "password", "", "The password for authentication")
	flag.StringVar(&subscription, "sub", "", "The URI of the subscription to delete")
	flag.Parse()

	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	service := c.Service

	es, err := service.EventService()
	if err != nil {
		panic(err)
	}

	err = es.DeleteEventSubscription(subscription)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted\n")
}
