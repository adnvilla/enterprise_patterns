package pedidos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ErrLockNotAcquired: otro usuario ya tiene reservado el recurso
var ErrLockNotAcquired = errors.New("el recurso está siendo editado por alguien más")

// Lock Manager: único punto de entrada para adquirir y liberar locks
type LockManager struct {
	db *sql.DB
}

func NewLockManager(db *sql.DB) *LockManager {
	return &LockManager{db: db}
}

// Acquire: intenta adquirir el lock exclusivo de escritura.
// La PRIMARY KEY de locks garantiza que solo un INSERT gana.
func (m *LockManager) Acquire(ctx context.Context, resource, owner string) error {
	res, err := m.db.ExecContext(ctx,
		`INSERT INTO locks (resource, owner, acquired_at)
		 VALUES ($1, $2, now())
		 ON CONFLICT (resource) DO NOTHING`,
		resource, owner)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// 0 filas insertadas = el lock ya era de alguien más
	if affected == 0 {
		return ErrLockNotAcquired
	}
	return nil
}

// Release: libera el lock SOLO si somos los dueños
func (m *LockManager) Release(ctx context.Context, resource, owner string) error {
	_, err := m.db.ExecContext(ctx,
		`DELETE FROM locks WHERE resource = $1 AND owner = $2`,
		resource, owner)
	return err
}

// OrderResource: nombre canónico del recurso a bloquear
func OrderResource(orderID int64) string {
	return fmt.Sprintf("order:%d", orderID)
}
