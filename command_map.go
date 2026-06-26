package main

import (
	"fmt"
)

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
