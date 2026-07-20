package postgres

import (
	"context"
	"database/sql"

	"github.com/adnvilla/enterprise_patterns/10-base/separated-interface/dominio" // la flecha de dependencia apunta AL dominio
)

// Implementación concreta de dominio.OrderRepository
type OrderRepository struct {
	db *sql.DB
}

// Gracias a las interfaces implícitas de Go, satisface el contrato
// del dominio sin declarar nada
var _ dominio.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) ByID(ctx context.Context, id int64) (*dominio.Order, error) {
	o := &dominio.Order{}
	err := r.db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents FROM orders WHERE id = $1`,
		id).Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *OrderRepository) Save(ctx context.Context, o *dominio.Order) error {
	if o.IsNew() { // cortesía del Layer Supertype
		return r.db.QueryRowContext(ctx,
			`INSERT INTO orders (customer_id, status, total_cents, version)
			 VALUES ($1, $2, $3, 1) RETURNING id`,
			o.CustomerID, o.Status, o.TotalCents).Scan(&o.ID)
	}
	_, err := r.db.ExecContext(ctx,
		`UPDATE orders SET status = $1, total_cents = $2, version = version + 1
		 WHERE id = $3`,
		o.Status, o.TotalCents, o.ID)
	return err
}
