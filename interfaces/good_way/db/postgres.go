package db

import (
	"database/sql"
	"fmt"
	"good-way/config"

	_ "github.com/lib/pq"
)

type Postgres struct {
	cfg config.DatabaseConfig
	db  *sql.DB
}

func NewPostgres(cfg config.DatabaseConfig) *Postgres {
	return &Postgres{cfg: cfg}
}

func (p *Postgres) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.cfg.Host,
		p.cfg.Port,
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) DB() any {
	return p.db
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
