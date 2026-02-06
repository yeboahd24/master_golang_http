package main

import "fmt"

type Calculator interface {
	Add(a, b int) int
	GetMemory() int
}

type CalculatorImpl struct {
	memory int
}

func (c *CalculatorImpl) Add(a, b int) int {
	c.memory += a + b
	return c.memory
}

func (c *CalculatorImpl) GetMemory() int {
	return c.memory
}

func main() {
	c := CalculatorImpl{memory: 0}
	c.Add(1, 2)
	c.Add(3, 4)
	fmt.Println(c.GetMemory())
}

// ============ STORAGE LAYER ============
// Interface for storage (returned from constructor)
type Storage interface {
	Save(key, value string) error
}

// Concrete implementation
type memoryStorage struct {
	data map[string]string
}

func (m *memoryStorage) Save(key, value string) error {
	m.data[key] = value
	return nil
}

// Returns INTERFACE - allows swapping S3, Redis, etc.
func NewStorage() Storage {
	return &memoryStorage{data: make(map[string]string)}
}

// ============ SERVICE LAYER ============
// Concrete struct (returned from constructor)
type UserService struct {
	store Storage // Depends on interface, not concrete type
}

// Returns STRUCT - specific to this service
func NewUserService(store Storage) *UserService {
	return &UserService{store: store}
}

func (s *UserService) CreateUser(name string) error {
	return s.store.Save("user:1", name)
}

// ============ MAIN ============
func main() {
	// Wire together: interface returned, struct accepted
	store := NewStorage()            // Returns Storage (interface)
	service := NewUserService(store) // Returns *UserService (struct)

	service.CreateUser("Alice")
	fmt.Println("User created")
}
