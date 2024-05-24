package rest

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/dashboard"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
	"github.com/tomassar/judicial-collection-case-management/internal/http/middleware"
)

type handler struct {
	cases     cases.Service
	users     users.Service
	auth      auth.Service
	lawyers   lawyers.Service
	dashboard dashboard.Service
}

func NewHandler(cases cases.Service, users users.Service, auth auth.Service, lawyers lawyers.Service, dashboard dashboard.Service) *handler {
	return &handler{
		cases:     cases,
		users:     users,
		auth:      auth,
		lawyers:   lawyers,
		dashboard: dashboard,
	}
}

// Init initializes the routes for the REST API
func (h *handler) Init() *gin.Engine {
	router := gin.Default()

	//middlewares
	injectUser := middleware.InjectUser(h.users, h.lawyers)
	requireAuth := middleware.RequireAuth(h.users, h.lawyers)

	router.Static("/static", "./static")

	// set cspmiddleware to all routes
	router.Use(middleware.CSPMiddleware())

	//home
	router.GET("/", injectUser, rootPath)
	//cases
	router.GET("/cases", requireAuth, getCases(h.cases))
	router.POST("/cases", injectUser, createCase(h.cases))
	router.DELETE("/cases/:id", requireAuth, deleteCase(h.cases))

	//users
	router.GET("/profiles/:id", getUserProfileByID(h.users))
	router.GET("/profiles/me", requireAuth, getUserProfile())

	//auth
	router.GET("/login", injectUser, getLogin())
	router.POST("/login", login(h.auth))
	router.GET("/signup", injectUser, getSignup())
	router.POST("/signup", signup(h.auth))
	router.POST("/logout", logout(h.auth))

	//dashboard
	router.GET("/dashboard", requireAuth, getDashboard(h.dashboard))

	return router
}

func rootPath(c *gin.Context) {
	//if user is already log in then redirect to /dashboard
	if u := getUserFromCtx(c); u != nil {
		slog.Info("User is already logged in", "user", u)
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	c.Redirect(http.StatusOK, "/login")
}

func getUserFromCtx(ctx context.Context) *users.User {
	user, ok := ctx.Value("user").(*users.User)
	slog.Info("getUserFromCtx slog", "user", user, "ok", ok)
	if !ok {
		return nil
	}

	return user
}
