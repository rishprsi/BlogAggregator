package main

import (
	"github.com/rishprsi/BlogAggregator/internal/config"
	"github.com/rishprsi/BlogAggregator/internal/database"
)

type state struct {
	db     *database.Queries
	Config *config.Config
}
