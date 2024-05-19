package handler

import (
	"goAuthService/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) parseToken(c *gin.Context) {
	var token dtos.TokenRequest
	if err := c.BindJSON(&token); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Auth.ParseToken(token.Token)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}