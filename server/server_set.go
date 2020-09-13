package server

import (
	"github.com/Nyarum/brrraaa/services"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(services.ServicesSet, wire.Struct(new(Server), "*"))
