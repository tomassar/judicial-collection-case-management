package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
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
	router.GET("/profiles/:id", getUserProfileByID(h.users))
	router.GET("/profiles/me", middleware.RequireAuth(h.users), getUserProfile())

	//auth
	router.POST("/login", login(h.auth))
	router.POST("/signup", signup(h.auth))
	return router
}

func getUserFromCtx(ctx context.Context) *users.User {
	return ctx.Value("user").(*users.User)
}
