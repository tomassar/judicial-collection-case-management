package cases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(group *gin.RouterGroup) {
	group.GET("/", getCases)
}

func getCases(c *gin.Context) {
	cases := GetCases()

	c.JSON(http.StatusOK, cases)
}
