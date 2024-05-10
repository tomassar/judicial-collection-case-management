package users

import (
	"context"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID uint) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserByID(ctx context.Context, userID uint) (*User, error) {
	user, err := s.repo.findByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Create(ctx context.Context, user *User) error {
	return s.repo.create(user)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repo.findByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
