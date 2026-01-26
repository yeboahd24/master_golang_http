package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, email string) error {
	return s.repo.Create(ctx, User{
		Name:  name,
		Email: email,
	})
}
