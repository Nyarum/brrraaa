package server

import (
	"github.com/Nyarum/brrraaa/core"
)

type Client struct {
	db *core.DB
}

func NewClient(db *core.DB) (*Client, error) {
	return &Client{}, nil
}
