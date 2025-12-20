package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rishprsi/BlogAggregator/internal/database"
	"github.com/rishprsi/BlogAggregator/internal/rss"
)

func handlerAggregator(s *state, cmd command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(ctx, url)
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("incorrect usage of addfeed, usage: addfeed 'name' 'url'")
	}
	ctx := context.Background()

	createdFeed, err := addFeedHelper(s, cmd, user, ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Feed with the name: %v and url: %v created and stored\n", createdFeed.Name, createdFeed.Url)
	return nil
}

func addFeedHelper(s *state, cmd command, user database.User, ctx context.Context) (database.Feed, error) {
	feed, err := s.db.GetFeed(ctx, cmd.args[1])
	if err == nil {
		fmt.Println("feed already exists")
	} else {
		now := sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		feedParams := database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      cmd.args[0],
			Url:       cmd.args[1],
			UserID:    user.ID,
		}
		fmt.Println("Creating the feed", feedParams.UserID)
		feed, err = s.db.CreateFeed(ctx, feedParams)
		if err != nil {
			return database.Feed{}, fmt.Errorf("error creating a feed: %v", err)
		}
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	createdFeedFollows, err := s.db.CreateFeedFollow(ctx, feedFollow)
	if err != nil {
		return database.Feed{}, fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Printf("Successfully created Feed: %v", createdFeedFollows.Feedname)

	return feed, nil
}

func handlerGetFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		if err != nil {
			fmt.Println("Failed to retreive user name")
		} else {
			fmt.Printf("Feed user: %v\n", user)
		}
	}

	return nil
}
