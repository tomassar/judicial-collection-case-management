package lawyers

import (
	"gorm.io/gorm"
)

type LawyerRepository interface {
}

type lawyerRepository struct {
	db *gorm.DB
}

func NewLawyerRepository(db *gorm.DB) LawyerRepository {
	return &lawyerRepository{
		db: db,
	}
}
