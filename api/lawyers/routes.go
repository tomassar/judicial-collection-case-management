package lawyers

import "github.com/gin-gonic/gin"

type lawyerRoutes struct {
	svc LawyerService
}

func NewUserRoutes(svc LawyerService) *lawyerRoutes {
	return &lawyerRoutes{
		svc: svc,
	}
}

func (c *lawyerRoutes) Init(router *gin.Engine) {
}
