package providers

import (
	"github.com/Nyarum/brrraaa/providers/config"
	"github.com/Nyarum/brrraaa/providers/database"
	"github.com/google/wire"
)

var ProvidersSet = wire.NewSet(config.NewConfig, database.NewDatabase)
