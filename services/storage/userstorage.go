package storage

import (
	"github.com/Nyarum/brrraaa/models"
	"github.com/Nyarum/brrraaa/providers/database"
)

type UserStorage struct {
	db database.Database
}

func NewUserStorage(db database.Database) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u UserStorage) Create(user, password string) (*models.User, error) {
	//models.NewUserStore()

	return nil, nil
}
