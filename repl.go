package main

import (
	"fmt"
	"os"
)

func RunRepl(s *state, cmds *commands) error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("invalid call: Usage: go run . command ...command_args")
	}

	cmd := command{
		name: args[1],
		args: []string{},
	}

	if len(args) > 2 {
		cmd.args = args[2:]
	}

	registerCommands(cmds)
	err := cmds.run(s, cmd)
	if err != nil {
		return fmt.Errorf("Call run failed:\n%v", err)
	}
	return nil
}

func registerCommands(cmds *commands) {
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerCreateUser)
	cmds.register("reset", handlerResetUsers)
	cmds.register("users", handlerListUsers)
}
