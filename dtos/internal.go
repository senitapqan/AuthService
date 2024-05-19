package dtos

import "goAuthService/models"

type TokenRequest struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	UserId int                   `json:"user_id"`
	Email  string                `json:"email"`
	Roles  []models.RolesHeaders `json:"roles"`
}
