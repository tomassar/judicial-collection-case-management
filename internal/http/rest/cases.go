package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/templates/cases_view"
	"github.com/tomassar/judicial-collection-case-management/internal/templates/toast"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
)

// getCases returns a handler for GET /cases requests
func getCases(s cases.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cases, err := s.GetLawyerCases(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get cases",
			})
			return
		}

		utils.RenderView(ctx, cases_view.List(cases))
	}
}

// create case returns a handler for POST /cases requests
func createCase(s cases.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body *cases.CreateCaseReq
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		lawyerID, err := utils.GetLawyerIDFromCtx(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		body.LawyerID = lawyerID
		err = s.CreateCase(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		toast.Success(ctx, "Successfully signed up to the newsletter!")
		ctx.Status(http.StatusOK)
	}
}

// delete case returns a handler for DELETE /cases/:id requests
func deleteCase(s cases.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = s.DeleteCase(ctx, uint(idInt))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(http.StatusOK)
	}
}
