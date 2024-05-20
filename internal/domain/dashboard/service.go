package dashboard

import (
	"context"

	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
)

type Service interface {
	GetDashboardData(ctx context.Context) (*DisplayDataRes, error)
}

type service struct {
	caseService cases.Service
}

func NewService(caseService cases.Service) Service {
	return &service{
		caseService: caseService,
	}
}

func (s *service) GetDashboardData(ctx context.Context) (*DisplayDataRes, error) {
	cs, err := s.caseService.GetLawyerCases(ctx)
	if err != nil {
		return nil, err
	}

	return &DisplayDataRes{
		Cases: cs,
	}, nil
}
