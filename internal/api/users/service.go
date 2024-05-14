package users

import (
	"context"
)

type Repository interface {
	Create(c *User) error
	FindByEmail(email string) (*User, error)
	FindByID(userID uint) (*User, error)
}

type Service interface {
	GetUserByID(ctx context.Context, userID uint) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByID(ctx context.Context, userID uint) (*User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Create(ctx context.Context, user *User) error {
	return s.repo.Create(user)
}

func (s *service) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
