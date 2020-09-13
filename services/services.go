package services

import (
	"github.com/Nyarum/brrraaa/providers"
	"github.com/Nyarum/brrraaa/services/storage"
	"github.com/google/wire"
)

var ServicesSet = wire.NewSet(providers.ProvidersSet, storage.NewUserStorage)
