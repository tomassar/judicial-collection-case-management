package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/internal/api/users"
	"github.com/tomassar/judicial-collection-case-management/internal/utils"
	"github.com/tomassar/judicial-collection-case-management/internal/views/user"
)

// getUserProfileByID returns a handler for GET /users/:id requests
func getUserProfileByID(s users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the uint id from the url
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}

		// get the user by id
		us, err := s.GetUserByID(c, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get user",
			})
			return
		}

		// return the user
		utils.RenderView(c, user.Show(us))
	}
}

// getUserProfile returns a handler for GET /users requests
func getUserProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the user from the context
		user := getUserFromCtx(ctx)

		// return the user
		ctx.JSON(http.StatusOK, user)
	}
}
