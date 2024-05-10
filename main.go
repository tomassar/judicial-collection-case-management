package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/api/auth"
	"github.com/tomassar/judicial-collection-case-management/api/cases"
	"github.com/tomassar/judicial-collection-case-management/api/users"
	"github.com/tomassar/judicial-collection-case-management/internal/database"
	"github.com/tomassar/judicial-collection-case-management/internal/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	dsn := os.Getenv("DB")
	db, err := database.Connect(dsn)
	if err != nil {
		panic("failed to connect database")
	}
	router := gin.Default()

	caseRepo := cases.NewCaseRepository(db)
	caseService := cases.NewCaseService(caseRepo)
	caseRoutes := cases.NewCaseRoutes(caseService)

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userRoutes := users.NewUserRoutes(userService)

	getUserByID := func(userID uint) (*users.User, error) {
		return userService.GetUserByID(&gin.Context{}, userID)
	}

	middleware := &cases.Middleware{
		Authorization: auth.RequireAuth(getUserByID),
	}

	caseRoutes.Init(router.Group("/cases"), middleware)
	userRoutes.Init(router)

	router.Run(os.Getenv("HOST_ADDR"))
}
