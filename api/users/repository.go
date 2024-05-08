package users

import "gorm.io/gorm"

type UserRepository interface {
	create(c *User) error
	findByEmail(email string) (*User, error)
	findByID(userID string) (*User, error)
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

func (r *userRepository) findByEmail(email string) (user *User, err error) {
	err = r.db.Debug().First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) findByID(userID string) (user *User, err error) {
	err = r.db.Debug().First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
