package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// JSON REQ/RES

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Helper func to send json responses
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Helper func to read json request
func readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576                                    // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // limiting the size of the requewst body
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	// ensure only one json object
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must contain a single JSON object")
	}
	return nil
}

// Example(like repository)
type InMemoryStore struct {
	mu     sync.RWMutex
	users  map[int]*User
	nextID int
}

// contructor
func NewInMemorySore() *InMemoryStore {
	return &InMemoryStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (s *InMemoryStore) Create(username, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		ID:        s.nextID,
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}
	s.users[s.nextID] = user
	s.nextID++
	return user, nil
}

func (s *InMemoryStore) GetByID(id int) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exist := s.users[id]
	if !exist {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *InMemoryStore) GetAll() []*User {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *InMemoryStore) Update(id int, username, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exist := s.users[id]
	if !exist {
		return nil, errors.New(("user not found"))
	}
	user.Username = username
	user.Email = email
	return user, nil
}

func (s *InMemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.users[id]; !exist {
		return errors.New("user not found")
	}
	delete(s.users, id)

	return nil
}

// Application struct pattern
type Application struct {
	store  *InMemoryStore
	logger *slog.Logger
}

// route
func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", app.handleUsers)
	// mux.HandleFunc("/api/users", app.handleUser)
	// mux.HandleFunc("/health", app.handleHealth)

	// wrap middleware
	return app.logRequest(app.recoverPanic(mux))
}

func (app *Application) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.listUsers(w, r)
	case http.MethodPost:
		app.createUser(w, r)
	default:
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
	}
}

func (app *Application) listUsers(w http.ResponseWriter, r *http.Request) {
	user := app.store.GetAll()
	writeJSON(w, http.StatusOK, user)
}

func (app *Application) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := readJSON(w, r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := app.store.Create(req.Username, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, http.StatusCreated, user)
}

// middleware
func (app *Application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		app.logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapped.statusCode,
			"duration", time.Since(start).String(),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Panic recovery middleware
func (app *Application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.logger.Error("panic recovered", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Request ID middleware
func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}
