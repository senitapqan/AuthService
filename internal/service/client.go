package service

import (
	"goAuthService/models"

	"github.com/rs/zerolog/log"
)

func (s *AuthService) CreateClient(client models.User) (int, error) {
	client.Password = s.hashPassword(client.Password)
	log.Info().Msg("service send request to repository: create client request")
	client_id, err := s.repos.CreateClient(client)
	return client_id, err
}