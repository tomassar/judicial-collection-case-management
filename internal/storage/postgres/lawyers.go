package postgres

import "gorm.io/gorm"

type lawersRepository struct {
	db *gorm.DB
}

func NewLawyersRepository(db *gorm.DB) *lawersRepository {
	return &lawersRepository{
		db: db,
	}
}
