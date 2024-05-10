package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	svc AuthService
}

func NewAuthRoutes(svc AuthService) *authRoutes {
	return &authRoutes{
		svc: svc,
	}
}

func (c *authRoutes) Init(router *gin.Engine) {
	router.POST("/signup", c.signUp)
	router.POST("/login", c.login)
}

func (c *authRoutes) signUp(ctx *gin.Context) {
	body := &SignUpReq{}

	if err := ctx.Bind(body); err != nil {
		slog.Error("failed to read body", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	err := c.svc.signUp(ctx, body)
	if err != nil {
		slog.Error("error while signing up", "error", err)
		// TODO: @(tomassar) Handle code errors appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *authRoutes) login(ctx *gin.Context) {
	body := &LoginReq{}

	if err := ctx.Bind(body); err != nil {
		slog.Error("failed to read body", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	err := c.svc.login(ctx, body)
	if err != nil {
		slog.Error("error while signing up", "error", err)
		// TODO: @(tomassar) Handle code errors appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
