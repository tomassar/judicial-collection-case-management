package cases

import "gorm.io/gorm"

type CaseRepository interface {
	FindAll() ([]Case, error)
}

type caseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) CaseRepository {
	return &caseRepository{
		db: db,
	}
}

func (c *caseRepository) FindAll() ([]Case, error) {
	var cases []Case
	err := c.db.Model(&Case{}).Find(&cases).Error
	if err != nil {
		return nil, err
	}

	return cases, nil
}
