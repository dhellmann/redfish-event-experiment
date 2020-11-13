package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"

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

	// chassis, err := service.Chassis()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, chass := range chassis {
	// 	fmt.Printf("Chassis: %#v\n\n", chass.ODataID)
	// 	therm, err := chass.Thermal()
	// 	if err != nil {
	// 		fmt.Printf("  failed to read thermal: %s\n\n", err)
	// 	} else {
	// 		fmt.Printf("  Thermal: %#v\n\n", therm)
	// 	}
	// 	power, err := chass.Power()
	// 	if err != nil {
	// 		fmt.Printf("  failed to read power: %s\n\n", err)
	// 	} else {
	// 		fmt.Printf("  Power: %#v\n\n", power)
	// 	}

	// 	comps, err := chass.ComputerSystems()
	// 	if err != nil {
	// 		fmt.Printf("  failed to read computer systems: %s\n\n", err)
	// 		continue
	// 	}
	// 	for _, comp := range comps {
	// 		fmt.Printf("  computer system: %s\n\n", comp.ID)
	// 	}
	// }

	es, err := service.EventService()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Event Types:\n")
	for _, et := range es.EventTypesForSubscription {
		fmt.Printf("  %s\n", et)
	}

	fmt.Printf("\nRegistry Prefixes:\n")
	for _, prefix := range es.RegistryPrefixes {
		fmt.Printf("  %s\n", prefix)
	}

	fmt.Printf("\nResource Types:\n")
	for _, rt := range es.ResourceTypes {
		fmt.Printf("  %s\n", rt)
	}

	fmt.Printf("\nSubscriptions:\n")
	subscriptions, err := es.GetEventSubscriptions()
	if err != nil {
		panic(err)
	}
	for _, sub := range subscriptions {
		fmt.Printf("  %s (%s)\n", sub.ODataID, sub.Destination)
	}
}
