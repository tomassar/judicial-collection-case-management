package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getDashboard reeturns a handler for GET / requests
// if user is not logged in, it is sent to /login page
func getDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the dashboard",
		})
	}
}
