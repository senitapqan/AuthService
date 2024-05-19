package handler

import (
	"goAuthService/dtos"
	"goAuthService/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) signIn(c *gin.Context) {
	var request dtos.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started generating token")

	token, err := h.service.Auth.GenerateToken(request.Login, request.Password)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) signUp(c *gin.Context) {
	var request models.User

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	
	log.Info().Msg("started handling create client request")
	id, err := h.service.Auth.CreateClient(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"new user was succesfully added with id": id,
	})
}