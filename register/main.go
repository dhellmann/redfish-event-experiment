package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"

	"github.com/dhellmann/redfish-event-experiment/config"
)

func main() {
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

	fmt.Printf("\nExisting Subscriptions:\n")
	subscriptions, err := es.GetEventSubscriptions()
	if err != nil {
		panic(err)
	}
	for _, sub := range subscriptions {
		fmt.Printf("  %s (%s)\n", sub.ODataID, sub.Destination)
	}

	headers := map[string]string{
		"X-arbitrary-header": "value",
	}

	oem := map[string]string{}

	subscriptionURI, err := es.CreateEventSubscription(
		"https://"+appConfig.Receiver.Endpoint+"/",
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
