package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
)

// getCases returns a handler for GET /cases requests
func getCases(s cases.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cases, err := s.GetCases(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get cases",
			})
			return
		}

		ctx.JSON(http.StatusOK, cases)
	}
}

// create case returns a handler for POST /cases requests
func createCase(s cases.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body *cases.CreateCaseReq
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		err := s.CreateCase(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create case",
			})
			return
		}

		ctx.Status(http.StatusOK)
	}
}
