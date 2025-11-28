package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("user login expects an argument 'username'")
	}

	err := s.Config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	return nil
}
