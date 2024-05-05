package cases

import "gorm.io/gorm"

type CaseRepository interface {
	findAll() ([]Case, error)
	create(c *Case) error
}

type caseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) CaseRepository {
	return &caseRepository{
		db: db,
	}
}

func (r *caseRepository) findAll() ([]Case, error) {
	var cases []Case
	err := r.db.Model(&Case{}).Find(&cases).Error
	if err != nil {
		return nil, err
	}

	return cases, nil
}

func (r *caseRepository) create(c *Case) error {
	return r.db.Create(c).Error
}
