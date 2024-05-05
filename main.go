package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/api/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/database"
)

func main() {
	dsn := "postgresql://postgres:password@postgresdb:5432/jccm"
	db, err := database.Connect(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	caseRepo := cases.NewCaseRepository(db)
	caseService := cases.NewCaseService(caseRepo)
	caseRoutes := cases.NewCaseRoutes(caseService)
	caseRoutes.Init(router.Group("/cases"))

	router.Run("0.0.0.0:8080")
}
