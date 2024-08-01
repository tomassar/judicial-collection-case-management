package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/sync_cases"
	"github.com/tomassar/judicial-collection-case-management/internal/templates/toast"
)

// sync returns a handler for POST /sync requests
func syncCase(d sync_cases.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := d.SyncCases(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		toast.Success(c, "Causa sincronizada!")
		c.Status(http.StatusOK)
	}
}
