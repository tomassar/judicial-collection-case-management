package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
)

// login returns a handler for POST /login requests
func login(s auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body *auth.LoginReq
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		err := s.Login(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to login",
			})
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// signup returns a handler for POST /signup requests
func signup(s auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body *auth.SignUpReq

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})

			return
		}

		err := s.SignUp(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to sign up",
			})
			return
		}

		ctx.Status(http.StatusOK)
	}
}
