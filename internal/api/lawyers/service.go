package lawyers

type Repository interface {
}

type Service interface {
}

type service struct {
	repo Repository
}

func NewLawyerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
