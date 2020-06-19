package server

import "github.com/google/wire"

var ServerSet = wire.NewSet(wire.Struct(new(Server), "*"), NewClient)
