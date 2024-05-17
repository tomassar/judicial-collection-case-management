package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
	"github.com/tomassar/judicial-collection-case-management/internal/views/authentication"
)

// login returns a handler for POST /login requests
func login(s auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body *auth.LoginReq
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		err := s.Login(c, body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to login",
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// login returns a handler for GET /login requests
func loginView() gin.HandlerFunc {
	return func(c *gin.Context) {
		//renders the login page located in views/login_templ.go
		utils.RenderView(c, authentication.LoginIndex())
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
