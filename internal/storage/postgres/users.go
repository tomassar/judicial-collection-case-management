package postgres

import (
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(c *users.User) error {
	return r.db.Debug().Create(c).Error
}

func (r *userRepository) FindByEmail(email string) (user *users.User, err error) {
	err = r.db.Debug().First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(userID uint) (user *users.User, err error) {
	err = r.db.Debug().First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
