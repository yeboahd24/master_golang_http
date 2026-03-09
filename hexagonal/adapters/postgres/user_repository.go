package postgres

import (
	"database/sql"

	"github.com/yeboahd24/hexagonal/internal/domain"
)

// Postgres specific implementation
// This is imkplementation of the port
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email) VALUES ($1,$2,$3)",
		user.ID, user.Name, user.Email,
	)

	return err
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id,name,email FROM users WHERE id=$1", id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	return &user, err
}
