package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomassar/judicial-collection-case-management/internal/database"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/dashboard"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/sync_cases"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
	"github.com/tomassar/judicial-collection-case-management/internal/http/rest"
	"github.com/tomassar/judicial-collection-case-management/internal/storage/postgres"
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

	fmt.Println("DB: ", os.Getenv("DB"))
	dsn := os.Getenv("DB")
	db, err := database.Connect(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	storage := postgres.NewStorage(db)

	caseService := cases.NewService(storage.Cases)
	userService := users.NewService(storage.Users)
	lawyerService := lawyers.NewLawyerService(storage.Lawyers)
	authService := auth.NewService(userService, lawyerService)
	dashboardService := dashboard.NewService(caseService)
	sync_cases := sync_cases.NewService()

	router := rest.
		NewHandler(caseService, userService, authService, lawyerService, dashboardService, sync_cases).
		Init()

	router.Run(os.Getenv("HOST_ADDR"))

}
