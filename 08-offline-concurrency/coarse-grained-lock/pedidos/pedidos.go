package pedidos

import (
	"context"
	"database/sql"
	"errors"
)

var ErrConflict = errors.New("conflicto de concurrencia: el pedido fue modificado por alguien más")

// Raíz del agregado: Order es dueño de sus líneas
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
	Version    int64 // Coarse-Grained Lock: UNA versión para TODO el agregado
	Lines      []OrderLine
}

// Las líneas no tienen versión propia: las cubre la de la raíz
type OrderLine struct {
	ProductID int64
	Quantity  int
}

// Repository: única puerta de entrada al agregado
type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// FindByID: carga el agregado completo (raíz + líneas) con la versión de la raíz
func (r *OrderRepository) FindByID(ctx context.Context, id int64) (*Order, error) {
	o := &Order{}
	err := r.db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents, version
		 FROM orders WHERE id = $1`, id).
		Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents, &o.Version)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx,
		`SELECT product_id, quantity
		 FROM order_lines WHERE order_id = $1 ORDER BY product_id`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var l OrderLine
		if err := rows.Scan(&l.ProductID, &l.Quantity); err != nil {
			return nil, err
		}
		o.Lines = append(o.Lines, l)
	}
	return o, rows.Err()
}

// Save persiste el agregado completo. Verifica e incrementa la versión
// de la raíz SIEMPRE, aunque solo hayas tocado una línea.
func (r *OrderRepository) Save(ctx context.Context, o *Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// El candado grueso: la versión de la raíz protege al grupo entero
	res, err := tx.ExecContext(ctx,
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
	if affected == 0 {
		return ErrConflict // el candado de la raíz protege al agregado entero
	}

	// Con el lock de la raíz asegurado, reescribimos las líneas
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM order_lines WHERE order_id = $1`, o.ID); err != nil {
		return err
	}
	for _, l := range o.Lines {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO order_lines (order_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			o.ID, l.ProductID, l.Quantity); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	o.Version++
	return nil
}
