package service

import (
	"goAuthService/internal/repository"
	"goAuthService/models"
)

type Auth interface {
	GenerateToken(username, password string) (string, error)

	CreateClient(client models.User) (int, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}