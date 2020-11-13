package main

import (
	"flag"
	"fmt"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
)

func main() {
	var receiver string
	config := gofish.ClientConfig{
		Insecure:  true,
		BasicAuth: true,
	}
	flag.StringVar(&config.Endpoint, "endpoint", "", "The URL of the BMC")
	flag.StringVar(&config.Username, "user", "", "The username for authentication")
	flag.StringVar(&config.Password, "password", "", "The password for authentication")
	flag.StringVar(&receiver, "receiver", ":9090", "The URL of the event receiver")
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

	fmt.Printf("\nExisting Subscriptions:\n")
	subscriptions, err := es.GetEventSubscriptions()
	if err != nil {
		panic(err)
	}
	for _, sub := range subscriptions {
		fmt.Printf("%#v\n\n", sub)
	}

	headers := map[string]string{
		"X-arbitrary-header": "value",
	}

	oem := map[string]string{}

	subscriptionURI, err := es.CreateEventSubscription(
		receiver,
		[]redfish.EventType{
			redfish.AlertEventType,
			redfish.ResourceAddedEventType,
			redfish.ResourceRemovedEventType,
			redfish.ResourceUpdatedEventType,
			redfish.StatusChangeEventType,
		},
		headers,
		oem,
		redfish.RedfishEventDestinationProtocol,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("subscription: %s\n", subscriptionURI)
}
