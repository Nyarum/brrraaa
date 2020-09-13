package database

import (
	"github.com/Nyarum/brrraaa/providers/config"

	"github.com/go-pg/pg"
)

type Database struct {
	DB *pg.DB
}

func NewDatabase(conf config.Config) Database {
	db := pg.Connect(&pg.Options{
		Addr:     conf.Database.Host,
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Database: conf.Database.Name,
	})
	defer db.Close()

	return Database{
		DB: db,
	}
}
