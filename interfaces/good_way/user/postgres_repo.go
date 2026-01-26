package user

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users(name, email) VALUES($1, $2)",
		u.Name, u.Email,
	)
	return err
}
