package postgres

import (
	"context"

	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"gorm.io/gorm"
)

type lawersRepository struct {
	db *gorm.DB
}

func NewLawyersRepository(db *gorm.DB) *lawersRepository {
	return &lawersRepository{
		db: db.Debug(),
	}
}

func (r *lawersRepository) Create(ctx context.Context, l *lawyers.Lawyer) error {
	return r.db.Create(l).Error
}

func (r *lawersRepository) FindByUserID(ctx context.Context, userID uint) (*lawyers.Lawyer, error) {
	var l lawyers.Lawyer
	if err := r.db.Where("user_id = ?", userID).First(&l).Error; err != nil {
		return nil, err
	}

	return &l, nil
}
