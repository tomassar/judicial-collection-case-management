package postgres

import (
	"github.com/tomassar/judicial-collection-case-management/internal/api/cases"
	"gorm.io/gorm"
)

type casesRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) *casesRepository {
	return &casesRepository{
		db: db,
	}
}

func (s *casesRepository) FindAll() ([]*cases.Case, error) {
	var ents []*cases.Case
	err := s.db.Model(&cases.Case{}).Find(&ents).Error
	if err != nil {
		return nil, err
	}

	return ents, nil
}

func (s *casesRepository) Create(c *cases.Case) error {
	return s.db.Create(&c).Error
}
