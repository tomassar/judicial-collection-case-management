package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/cases"
	"gorm.io/gorm"
)

func main() {
	/* dsn := "postgresql://postgres:password@localhost:5432/jccm"
	db, err := database.Connect(dsn)
	if err != nil {
		panic("failed to connect database")
	} */

	router := gin.Default()

	caseRepo := cases.NewCaseRepository(&gorm.DB{})
	caseService := cases.NewCaseService(caseRepo)
	caseController := cases.NewCaseController(caseService)
	caseController.InitRoutes(router.Group("/cases"))

	router.Run(":8080")
}
