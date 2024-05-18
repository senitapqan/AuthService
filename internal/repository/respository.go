package repository

import (
	"goAuthService/models"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateClient() (int, error)
	GetUser(login, password string) (models.User, error) 
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}