package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/cases"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	cases.InitRoutes(router.Group("/cases"))

	router.Run(":8080")
}
