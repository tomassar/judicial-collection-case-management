package lawyers

import "context"

type Repository interface {
	Create(ctx context.Context, l *Lawyer) error
	FindByUserID(ctx context.Context, userID uint) (*Lawyer, error)
}

type Service interface {
	CreateLawyer(ctx context.Context, body *CreateLawyerReq) error
	GetLawyerByUserID(ctx context.Context, userID uint) (*Lawyer, error)
}

type service struct {
	repo Repository
}

func NewLawyerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateLawyer(ctx context.Context, body *CreateLawyerReq) error {
	return s.repo.Create(ctx, &Lawyer{
		UserID: body.UserID,
	})
}

func (s *service) GetLawyerByUserID(ctx context.Context, userID uint) (*Lawyer, error) {
	return s.repo.FindByUserID(ctx, userID)
}
