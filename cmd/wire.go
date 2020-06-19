//go:generate wire
//+build wireinject

package cmd

import (
	"github.com/Nyarum/brrraaa/core"
	"github.com/Nyarum/brrraaa/server"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	core.CoreSet,
	server.ServerSet,
)

func NewServer() (*server.Server, error) {
	panic(
		wire.Build(
			providerSet,
		),
	)
}
