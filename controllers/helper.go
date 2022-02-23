package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nemonicgod/terraforms-api/backend"
)

func GetBackend(c *gin.Context) (*backend.Backend, error) {
	b, err := c.Keys["backend"].(*backend.Backend)
	if !err {
		return nil, errors.New("could not get 'backend' context connection from gin.Context")
	}

	return b, nil
}
