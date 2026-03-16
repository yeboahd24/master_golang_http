package main

// D — Dependency Inversion Principle
// High-level modules shouldn't depend on low-level ones — both should depend on abstractions.
// In Go, the key insight is that interfaces belong to the consumer, not the producer.


// BAD — OrderProcessor depends directly on the concrete type
type MySQLOrderRepo struct{ db *sql.DB }
func (r *MySQLOrderRepo) Save(o Order) error { ... }

type OrderProcessor struct {
    repo *MySQLOrderRepo // tightly coupled to MySQL
}

// GOOD — OrderProcessor owns the interface it needs
type OrderProcessor struct {
    repo OrderSaver // depends on abstraction
}

type OrderSaver interface {
    Save(o Order) error
}

// MySQLOrderRepo satisfies it implicitly — no import of the interface needed
type MySQLOrderRepo struct{ db *sql.DB }
func (r *MySQLOrderRepo) Save(o Order) error { ... }

// So does a mock for testing
type MockOrderRepo struct{ saved []Order }
func (m *MockOrderRepo) Save(o Order) error {
    m.saved = append(m.saved, o)
    return nil
}
