package handler

import (
	"goAuthService/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   *service.Service
	validator RequestValidator
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:   service,
		validator: NewReuqestValidator(),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	internal := router.Group("/internal-api")
	{
		internal.POST("/parse-token", h.parseToken)
	}

	return router
}
