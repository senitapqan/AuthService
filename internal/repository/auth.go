package repository

import (
	"goAuthService/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateClient() (int, error) {
	return 0, nil
}

func (r *AuthRepository) GetUser(login, password string) (models.User, error) {
	return models.User{}, nil
}

