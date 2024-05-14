package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/api/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/api/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/api/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/api/users"
	"github.com/tomassar/judicial-collection-case-management/internal/http/middleware"
)

type handler struct {
	cases   cases.Service
	users   users.Service
	auth    auth.Service
	lawyers lawyers.Service
}

func NewHandler(cases cases.Service, users users.Service, auth auth.Service, lawyers lawyers.Service) *handler {
	return &handler{
		cases:   cases,
		users:   users,
		auth:    auth,
		lawyers: lawyers,
	}
}

// Init initializes the routes for the REST API
func (h *handler) Init() *gin.Engine {
	router := gin.Default()

	//cases
	router.GET("/cases", middleware.RequireAuth(h.users), getCases(h.cases))
	router.POST("/cases", createCase(h.cases))

	//users
	router.GET("/users/:id", getUserProfileByID(h.users))
	router.GET("/users/me", middleware.RequireAuth(h.users), getUserProfile())

	//auth
	router.POST("/login", login(h.auth))
	router.POST("/signup", signup(h.auth))
	return router
}

func getUserFromCtx(ctx context.Context) *users.User {
	return ctx.Value("user").(*users.User)
}
