package repository

import (
	"goAuthService/models"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	GetUser(login string) (models.User, error) 
	GetUserById(id int) (models.User, error)
	GetRoles(userId int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateClient(client models.User) (int, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}