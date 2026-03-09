package application

import (
	"github.com/yeboahd24/hexagonal/internal/domain"
	"github.com/yeboahd24/hexagonal/internal/ports"
)

// UserService-->UserRepository(interface)
// This is dependency inversion

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
	user := domain.User{
		ID:    "generated-id",
		Name:  name,
		Email: email,
	}

	return s.repo.Save(user)
}
