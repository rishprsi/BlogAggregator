package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to create request with the error: %v", err)
	}
	request.Header.Set("User-Agent", "gator")
	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to make the request: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed decoding xml: %v", err)
	}
	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed decoding xml: %v", err)
	}
	parseRSSFeed(&feed)
	return &feed, nil
}

func parseRSSFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Link = html.UnescapeString(feed.Channel.Link)
	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Link = html.UnescapeString(item.Link)
	}
}
