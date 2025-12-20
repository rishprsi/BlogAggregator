package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func createUUID() uuid.UUID {
	return uuid.New()
}

func createNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

func createTime() time.Time {
	return time.Now()
}
