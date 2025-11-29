package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rishprsi/BlogAggregator/internal/config"
	"github.com/rishprsi/BlogAggregator/internal/database"
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

	db, err := sql.Open("postgres", currState.Config.DBUrl)
	dbQueries := database.New(db)

	currState.db = dbQueries

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
