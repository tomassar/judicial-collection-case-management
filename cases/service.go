package cases

import "time"

type CaseService interface {
	GetCases() ([]Case, error)
}
type caseService struct {
	repo CaseRepository
}

func NewCaseService(repo CaseRepository) CaseService {
	return &caseService{
		repo: repo,
	}
}

func (s *caseService) GetCases() ([]Case, error) {
	cases := []Case{
		{ID: 1, DebtorName: "Victoria Contreras", Amount: 100050, CreationDate: time.Now(), Status: "In progress", Documents: []string{"Contract", "Invoices"}},
		{ID: 2, DebtorName: "Jane Smith", Amount: 200075, CreationDate: time.Now(), Status: "Pending", Documents: []string{"Agreement", "Receipts"}},
	}

	return cases, nil
}
