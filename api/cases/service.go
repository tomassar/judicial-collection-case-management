package cases

import (
	"context"
)

type CaseService interface {
	GetCases(ctx context.Context) ([]Case, error)
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

func (s *caseService) GetCases(ctx context.Context) ([]Case, error) {
	return s.repo.findAll()
}

func (s *caseService) CreateCase(ctx context.Context, body *CreateCaseReq) error {
	c := &Case{
		DebtorName: body.DebtorName,
		Status:     body.Status,
	}

	return s.repo.create(c)
}
