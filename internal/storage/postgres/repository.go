package postgres

import (
	"gorm.io/gorm"
)

type Storage struct {
	Cases   *casesRepository
	Lawyers *lawersRepository
	Users   *userRepository
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		Cases:   NewCaseRepository(db),
		Lawyers: NewLawyersRepository(db),
		Users:   NewUserRepository(db),
	}
}
