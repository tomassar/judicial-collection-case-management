package cases

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomassar/judicial-collection-case-management/api/users"
)

type CaseService interface {
	GetCases(c *gin.Context) ([]*caseEntity, error)
	CreateCase(ctx context.Context, body *CreateCaseReq) error
}
type caseService struct {
	repo CaseRepository
}

func NewCaseService(repo CaseRepository) CaseService {
	return &caseService{
		repo: repo,
	}
}

func (s *caseService) GetCases(c *gin.Context) ([]*caseEntity, error) {
	userValue, exists := c.Get("user")
	if !exists {
		slog.Error("user data not found in context")
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, nil
	}

	user, ok := userValue.(*users.User)
	if !ok {
		slog.Error("user data has unexpected type")
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, nil
	}

	slog.Info("got user", "user", user)

	return s.repo.findAll()
}

func (s *caseService) CreateCase(ctx context.Context, body *CreateCaseReq) error {
	c := &caseEntity{
		DebtorName: body.DebtorName,
		Status:     body.Status,
	}

	return s.repo.create(c)
}
