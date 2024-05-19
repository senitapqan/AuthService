package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) localOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Host != "localhost:8000" && c.Request.Host != "127.0.0.1:8000" {
			newErrorResponse(c, http.StatusForbidden, "Request from another service")
			return
		}
		c.Next()
	}
}