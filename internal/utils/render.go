package utils

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderView(c *gin.Context, component templ.Component) {
	component.Render(c.Request.Context(), c.Writer)
}
