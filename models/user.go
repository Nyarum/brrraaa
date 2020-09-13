package models

//go:generate kallax gen

import (
	"time"

	"gopkg.in/src-d/go-kallax.v1"
)

type User struct {
	kallax.Model `table:"users" pk:"id"`
	ID           kallax.ULID
	User         string
	Password     string
	Created      time.Time
	Updated      time.Time
}
