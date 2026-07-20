package sesion

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

// El estado de una interacción en curso: igual que en los otros
// patrones del bloque, solo cambia dónde vive.
type Session struct {
	CustomerID int64
	CartItems  []int64 // ids de productos
	ExpiresAt  time.Time
}

// Database Session State: la sesión es una fila más
type DBStore struct {
	db *sql.DB
}

func NewDBStore(db *sql.DB) *DBStore {
	return &DBStore{db: db}
}

func (s *DBStore) Save(ctx context.Context, id string, sess Session, ttl time.Duration) error {
	data, err := json.Marshal(sess)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO sessions (id, data, expires_at) VALUES ($1, $2, $3)
		 ON CONFLICT (id) DO UPDATE SET data = $2, expires_at = $3`,
		id, data, time.Now().Add(ttl))
	return err
}

func (s *DBStore) Load(ctx context.Context, id string) (Session, error) {
	var raw []byte
	err := s.db.QueryRowContext(ctx,
		`SELECT data FROM sessions WHERE id = $1 AND expires_at > now()`,
		id).Scan(&raw)
	if err != nil {
		return Session{}, err
	}
	var sess Session
	err = json.Unmarshal(raw, &sess)
	return sess, err
}
