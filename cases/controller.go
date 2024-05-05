package cases

import (
	"log/slog"
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
	group.POST("/", c.createCase)
}

func (c *caseController) getCases(ctx *gin.Context) {
	cases, err := c.svc.GetCases(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cases)
}

func (c *caseController) createCase(ctx *gin.Context) {
	var body *CreateCaseReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		slog.Error("error while decoding body", "error", err)
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := c.svc.CreateCase(ctx, body)
	if err != nil {
		slog.Error("error while creating case", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
