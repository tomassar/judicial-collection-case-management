package cases

import (
	"context"
)

type Repository interface {
	FindAll() ([]*Case, error)
	Create(c *Case) error
}

type Service interface {
	GetCases(ctx context.Context) ([]*Case, error)
	CreateCase(ctx context.Context, body *CreateCaseReq) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCases(ctx context.Context) ([]*Case, error) {
	return s.repo.FindAll()
}

func (s *service) CreateCase(ctx context.Context, body *CreateCaseReq) error {
	c := &Case{
		DebtorName: body.DebtorName,
		Status:     body.Status,
	}

	return s.repo.Create(c)
}
