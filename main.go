package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/api/cases"
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
	caseRoutes.Init(router.Group("/cases"))

	router.Run(os.Getenv("HOST_ADDR"))
}
