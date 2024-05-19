package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestValidator interface {
	ValidateId(c *gin.Context) (int, error) 
	ValidatePage(c *gin.Context) (int, error) 
}

type requestValidator struct {}

func NewReuqestValidator() RequestValidator {
	return &requestValidator{}
}

func (v *requestValidator) ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1, err
	}

	if id <= 0 {
		return -1, errors.New("id need to be positive")
	}
	return id, nil
}


func (v *requestValidator) ValidatePage(c *gin.Context) (int, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		return -1, errors.New("value of query parameter 'page' is not integer")
	}

	if page <= 0 {
		return -1, errors.New("page need to be positive")
	}

	return page, nil
}
