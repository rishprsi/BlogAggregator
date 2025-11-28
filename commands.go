package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}

	err := command(s, cmd)
	if err != nil {
		return fmt.Errorf("function invocation failed:\n%v", err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
