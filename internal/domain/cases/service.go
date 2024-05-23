package cases

import (
	"context"

	"github.com/tomassar/judicial-collection-case-management/internal/utils"
)

type Repository interface {
	FindAllByLawyerID(ctx context.Context, lawyerID uint) ([]*Case, error)
	FindAll() ([]*Case, error)
	Create(c *Case) error
	Delete(ctx context.Context, id uint) error
}

type Service interface {
	GetLawyerCases(ctx context.Context) ([]*Case, error)
	CreateCase(ctx context.Context, body *CreateCaseReq) error
	DeleteCase(ctx context.Context, id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type GetLawyerCasesReq struct {
	LawyerID string `json:"lawyer_id"`
}

func (s *service) GetLawyerCases(ctx context.Context) ([]*Case, error) {
	lawyerID, err := utils.GetLawyerIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return s.repo.FindAllByLawyerID(ctx, lawyerID)
}

func (s *service) CreateCase(ctx context.Context, body *CreateCaseReq) error {
	c := &Case{
		DebtorName: body.DebtorName,
		Status:     body.Status,
		Role:       body.Role,
		Date:       body.Date.ToTime(),
		Subject:    body.Subject,
		Court:      body.Court,
		LawyerID:   body.LawyerID,
	}

	return s.repo.Create(c)
}

func (s *service) DeleteCase(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
