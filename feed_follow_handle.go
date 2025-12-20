package main

import (
	"context"
	"fmt"

	"github.com/rishprsi/BlogAggregator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.args) < 1 {
		return fmt.Errorf("incorrect usage of the command follow Usage: follow 'url'")
	}
	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed not found in the database: %v", err)
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

func handlerGetUserFeeds(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return fmt.Errorf("failed to get user feeds for user: %v with error: %v", s.Config.CurrentUserName, err)
	}
	fmt.Printf("The user %v is following the below feeds:\n", s.Config.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.Feedname)
	}
	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("incorrect usage of the command follow Usage: unfollow 'feed_id'")
	}
	ctx := context.Background()
	params := database.UnfollowFeedParams{
		Name: user.Name,
		Url:  cmd.args[0],
	}
	err := s.db.UnfollowFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to get delete feeds for user: %v with error: %v", s.Config.CurrentUserName, err)
	}
	return nil
}
