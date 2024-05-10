package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tomassar/judicial-collection-case-management/api/auth"
	"github.com/tomassar/judicial-collection-case-management/api/cases"
	"github.com/tomassar/judicial-collection-case-management/api/users"
	"github.com/tomassar/judicial-collection-case-management/internal/database"
)

func main() {
	initService()
}

func initService() {
	//init env variables with godotenv
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env variables")
	}

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

	authService := auth.NewAuthService(userService)
	authRoutes := auth.NewAuthRoutes(authService)

	middleware := cases.Middleware{
		Authorization: authService.RequireAuth,
	}
	caseRoutes.Init(router.Group("/cases"), &middleware)
	userRoutes.Init(router)
	authRoutes.Init(router)

	router.Run(os.Getenv("HOST_ADDR"))
}
