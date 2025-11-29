package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rishprsi/BlogAggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("user login expects an argument 'username'")
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("User with the username: %v does not exist in the database", cmd.args[0])
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	return nil
}

func handlerCreateUser(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("user register expects and argument 'name'")
	}
	ctx := context.Background()
	existingUser, err := s.db.GetUser(ctx, cmd.args[0])

	if err == nil {
		return fmt.Errorf("Found an existing user with the given name: %v", existingUser.Name)
	}
	currTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      cmd.args[0],
	}

	createdUser, err := s.db.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("Error creating user got: %v", err)
	}
	s.Config.SetUser(createdUser.Name)
	fmt.Println("The current user is now set to: ", createdUser.Name)
	return nil
}

func handlerResetUsers(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.ResetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Failed to delete all users: %v", err)
	}
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get users: %v", err)
	}
	for _, user := range users {
		conditional := ""
		if s.Config.CurrentUserName == user.Name {
			conditional = " (current)"
		}
		fmt.Printf("* %v%v\n", user.Name, conditional)
	}
	return nil
}
