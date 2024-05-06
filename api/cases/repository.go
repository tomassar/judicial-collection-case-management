package cases

import "gorm.io/gorm"

type CaseRepository interface {
	findAll() ([]*caseEntity, error)
	create(c *caseEntity) error
}

type caseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) CaseRepository {
	return &caseRepository{
		db: db,
	}
}

func (r *caseRepository) findAll() ([]*caseEntity, error) {
	var cases []*Case
	err := r.db.Model(&Case{}).Find(&cases).Error
	if err != nil {
		return nil, err
	}

	return caseModelsToEntities(cases), nil
}

func (r *caseRepository) create(c *caseEntity) error {
	return r.db.Create(c.toModel()).Error
}
