package db

// Depend on behavior not implementation
type Database interface {
	Connect() error
	DB() any
	Close() error
}
