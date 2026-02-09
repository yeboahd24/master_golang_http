package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

var ErrSessionExists = newError("session exists")

func generateSessionID() string {
	return "session-" + randomString(16)
}

// randomString generates a random string of the given length
func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// newError returns a new error with the given message
func newError(msg string) error {
	return errors.New(msg)
}

// ID is the session ID
// UserID is the user ID
// Data is the session data
// ExpireAt is the expire time
type session struct {
	ID       string
	UserID   string
	Data     map[string]any
	ExpireAt time.Time
}

// Store is the session store

type Store struct {
	mux      sync.RWMutex
	sessions map[string]*session
	ttl      time.Duration
}

func NewStore(ttl time.Duration) *Store {
	s := &Store{
		sessions: make(map[string]*session),
		ttl:      ttl,
	}
	go s.cleanup()
	return s
}

func (s *Store) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		s.mux.Lock()
		now := time.Now()
		for k, v := range s.sessions {
			if now.After(v.ExpireAt) {
				delete(s.sessions, k)
			}
		}
		s.mux.Unlock()
	}
}

func (s *Store) Create(userID string) (*session, error) {
	sessionID := generateSessionID()
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, ok := s.sessions[sessionID]; ok {
		return nil, ErrSessionExists
	}
	s.sessions[sessionID] = &session{
		ID:       sessionID,
		UserID:   userID,
		Data:     make(map[string]any),
		ExpireAt: time.Now().Add(s.ttl),
	}
	return s.sessions[sessionID], nil
}

func (s *Store) Get(sessionID string) (*session, bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	session, ok := s.sessions[sessionID]
	if !ok || time.Now().After(session.ExpireAt) {
		return nil, false
	}
	return session, true
}

func (s *Store) Delete(sessionID string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.sessions, sessionID)
}
