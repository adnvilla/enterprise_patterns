package pedidos

import (
	"context"
	"database/sql"
	"errors"
)

// ErrConflict: la versión cambió desde que leímos el pedido
var ErrConflict = errors.New("conflicto de concurrencia: el pedido fue modificado por alguien más")

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
	Version    int64 // la huella de qué versión leímos
}

// Repository con Optimistic Offline Lock
type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// FindByID: lee el pedido INCLUYENDO su versión
func (r *OrderRepository) FindByID(ctx context.Context, id int64) (*Order, error) {
	o := &Order{}
	err := r.db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents, version
		 FROM orders WHERE id = $1`, id).
		Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents, &o.Version)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Update: solo escribe si la versión sigue siendo la que leímos
func (r *OrderRepository) Update(ctx context.Context, o *Order) error {
	res, err := r.db.ExecContext(ctx,
		`UPDATE orders
		 SET status = $1, total_cents = $2, version = version + 1
		 WHERE id = $3 AND version = $4`,
		o.Status, o.TotalCents, o.ID, o.Version)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// 0 filas afectadas = alguien se nos adelantó
	if affected == 0 {
		return ErrConflict
	}
	o.Version++ // reflejamos en memoria la nueva versión
	return nil
}
