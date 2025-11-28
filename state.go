package main

import (
	"database"

	"github.com/rishprsi/BlogAggregator/internal/config"
)

type state struct {
	db     *database.Queries
	Config *config.Config
}
