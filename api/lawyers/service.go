package lawyers

type LawyerService interface {
}

type lawyerService struct {
	repo LawyerRepository
}

func NewLawyerService(repo LawyerRepository) LawyerService {
	return &lawyerService{
		repo: repo,
	}
}
