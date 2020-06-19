package core

import "github.com/google/wire"

var CoreSet = wire.NewSet(NewConfig, NewDB)
