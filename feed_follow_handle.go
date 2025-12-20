package main

import (
	"context"
	"fmt"

	"github.com/rishprsi/BlogAggregator/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) < 1 {
		return fmt.Errorf("incorrect usage of the command follow Usage: follow 'url'")
	}
	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed not found in the database: %v", err)
	}
	user, err := s.db.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Current user not found in the database: %v", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        createUUID(),
		CreatedAt: createTime(),
		UpdatedAt: createTime(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("failed to create Feed Follows: %v", err)
	}
	fmt.Printf("%v successfully followed the feed: %v", feedFollow.Username, feedFollow.Feedname)

	return nil
}

func handlerGetUserFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeedFollowsForUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user feeds for user: %v with error: %v", s.Config.CurrentUserName, err)
	}
	fmt.Printf("The user %v is following the below feeds:\n", s.Config.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.Feedname)
	}
	return nil
}
