package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	return nil
}
