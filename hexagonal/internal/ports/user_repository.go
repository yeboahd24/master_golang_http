package ports

import "github.com/yeboahd24/hexagonal/internal/domain"

// The application depends on this interface
type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (*domain.User, error)
}
