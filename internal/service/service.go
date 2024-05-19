package service

import (
	"goAuthService/dtos"
	"goAuthService/internal/repository"
	"goAuthService/models"
)


type Auth interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (dtos.TokenResponse, error)

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