package cases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type caseController struct {
	svc CaseService
}

func NewCaseController(svc CaseService) *caseController {
	return &caseController{
		svc: svc,
	}
}

func (c *caseController) InitRoutes(group *gin.RouterGroup) {
	group.GET("/", c.getCases)
}

func (c *caseController) getCases(ctx *gin.Context) {
	cases, err := c.svc.GetCases()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cases)
}
