package main

import (
	"fmt"
	"os"

	"github.com/rishprsi/BlogAggregator/internal/config"
)

func main() {
	fmt.Println("DB, Go Project")

	configStruct, err := readConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	currState := state{
		Config: &configStruct,
	}

	currCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	err = RunRepl(&currState, &currCommands)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readConfig() (config.Config, error) {
	configStruct, err := config.Read()
	if err != nil {
		return config.Config{}, err
	}
	return configStruct, nil
}
