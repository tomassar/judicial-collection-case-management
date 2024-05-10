package cases

import (
	"context"

	"github.com/tomassar/judicial-collection-case-management/api/users"
)

type CaseService interface {
	GetCases(ctx context.Context) ([]*caseEntity, error)
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

func (s *caseService) GetCases(ctx context.Context) ([]*caseEntity, error) {
	return s.repo.findAll()
}

func (s *caseService) CreateCase(ctx context.Context, body *CreateCaseReq) error {
	c := &caseEntity{
		DebtorName: body.DebtorName,
		Status:     body.Status,
	}

	return s.repo.create(c)
}

func getUserFromCtx(ctx context.Context) *users.User {
	return ctx.Value("user").(*users.User)
}
