package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
	"github.com/tomassar/judicial-collection-case-management/internal/views/dashboard"
)

// getDashboard reeturns a handler for GET / requests
// if user is not logged in, it is sent to /login page
func getDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.RenderView(c, dashboard.Show())
	}
}
