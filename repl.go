package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rishprsi/BlogAggregator/internal/database"
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
		return fmt.Errorf("call run failed:\n%v", err)
	}

	return nil
}

func registerCommands(cmds *commands) {
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerCreateUser)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAggregator)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFeedFollow))
	cmds.register("following", middlewareLoggedIn(handlerGetUserFeeds))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
