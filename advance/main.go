package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// database
type DBStore struct {
	db *sql.DB
}

func NewDBStore(connStr string) (*DBStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DBStore{db: db}, nil
}

func (s *DBStore) CreateUser(ctx context.Context, username, email, hashedPassword string) (*User, error) {
	query := `
		INSERT INTO users (username, email, password_hash, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, username, email, created_at
	`

	var user User
	err := s.db.QueryRowContext(ctx, query, username, email, hashedPassword).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// graceful shutdown
func runServerWithGracefullShutdown() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	store := NewDBStore("postgres://username@password")

	app := &Application{
		store:  store,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start server in goroutine
	go func() {
		logger.Info("starting server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}
	logger.Info("server stopped")
}
