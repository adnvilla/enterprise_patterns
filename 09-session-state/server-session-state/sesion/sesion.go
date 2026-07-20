package sesion

import (
	"sync"
	"time"
)

// Server Session State: los datos viven aquí, el cliente solo lleva el id
type Session struct {
	CustomerID int64
	CartItems  []int64 // ids de productos
	ExpiresAt  time.Time
}

type Store struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func NewStore() *Store {
	return &Store{sessions: make(map[string]Session)}
}

func (s *Store) Put(id string, sess Session, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess.ExpiresAt = time.Now().Add(ttl)
	s.sessions[id] = sess
}

func (s *Store) Get(id string) (Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	if !ok || time.Now().After(sess.ExpiresAt) {
		return Session{}, false
	}
	return sess, true
}
