package users

import (
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	svc UserService
}

func NewUserRoutes(svc UserService) *userRoutes {
	return &userRoutes{
		svc: svc,
	}
}

func (c *userRoutes) Init(router *gin.Engine) {
}
