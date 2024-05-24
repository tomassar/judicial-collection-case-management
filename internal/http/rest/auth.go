package rest

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/templates/authentication"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
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

		c.Header("HX-Redirect", "/dashboard")
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	}
}

// login returns a handler for GET /login requests
func getLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if user is already log in then redirect to /dashboard
		if u := getUserFromCtx(c); u != nil {
			slog.Info("User is already logged in", "user", u)
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		utils.RenderView(c, authentication.LoginIndex())
	}
}

// getSignup returns a handler for GET /signup requests
func getSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if user is already log in then redirect to /dashboard
		if u := getUserFromCtx(c); u != nil {
			slog.Info("User is already logged in", "user", u)
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		utils.RenderView(c, authentication.SignUp())
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

// logout returns a handler for POST /logout requests
func logout(s auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		s.Logout(c)

		c.Header("HX-Redirect", "/login")
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful",
		})
	}
}
