package service

import (
	"context"
	"errors"
	"myapp/model"
	"myapp/repository"
)

type UserService interface {
	Register(ctx context.Context, email, name string) (*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository // depends on interface!
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, email, name string) (*model.User, error) {
	// Business logic: validation, hashing, etc.
	if !isValidEmail(email) {
		return nil, errors.New("invalid email")
	}

	user := &model.User{
		ID:    generateID(),
		Email: email,
		Name:  name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
