package users

import "gorm.io/gorm"

type UserRepository interface {
	create(c *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) create(c *User) error {
	return r.db.Debug().Create(c).Error
}
