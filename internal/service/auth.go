package service

import "goAuthService/internal/repository"

type AuthService struct {
	repos repository.Auth
}

func NewAuthService(repos repository.Auth) *AuthService {
	return &AuthService{
		repos: repos,
	}
}