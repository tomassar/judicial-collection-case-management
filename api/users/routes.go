package users

import (
	"log/slog"
	"net/http"

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
	router.POST("/signup", c.signUp)
}

func (c *userRoutes) signUp(ctx *gin.Context) {
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
