package repository

import (
	"context"
	"database/sql"
	"myapp/model"
)

// Interface - makes testing easy!
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

// Implementation
type userRepo struct {
	db *sql.DB
}

// Constructor returns the interface, not concrete type!
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	// SQL implementation...
}
