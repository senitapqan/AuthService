package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"goAuthService/dtos"
	"goAuthService/internal/repository"
	"goAuthService/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repos repository.Auth
}

func NewAuthService(repos repository.Auth) *AuthService {
	return &AuthService{
		repos: repos,
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	Roles  []models.RolesHeaders
}

func (s *AuthService) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	password = s.hashPassword(password)
	log.Info().Msg("service send request to repository: get user request")
	user, err := s.repos.GetUser(login)

	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("Incorrect password")
	}

	var rolesHeaders []models.RolesHeaders
	log.Info().Msg("service send request to repository: get roles request")
	roles, err := s.repos.GetRoles(user.Id)

	if err != nil {
		return "", err
	}

	for _, role := range roles {
		id, err := s.repos.GetRoleId(role, user.Id)
		if err != nil {
			return "", err
		}
		rolesHeaders = append(rolesHeaders, models.RolesHeaders{Role: role, Id: id})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		rolesHeaders,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	return tokenString, err
}

func (s *AuthService) ParseToken(accessToken string) (dtos.TokenResponse, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return dtos.TokenResponse{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return dtos.TokenResponse{}, errors.New("token claims are not type of *tokenClaims")
	}

	user, err := s.repos.GetUserById(claims.UserId)

	if err != nil {
		return dtos.TokenResponse{}, err
	}

	return dtos.TokenResponse{
			UserId: claims.UserId, 
			Roles: claims.Roles,
			Email: user.Email,
		}, nil
}