package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/dashboard"
	"github.com/tomassar/judicial-collection-case-management/internal/templates/dashboard_view"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
)

// getDashboard reeturns a handler for GET / requests
// if user is not logged in, it is sent to /login page
func getDashboard(d dashboard.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		displayData, err := d.GetDashboardData(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		utils.RenderView(c, dashboard_view.Show(displayData.Cases))
	}
}
