package pedidos

import (
	"context"
	"database/sql"
	"sync"
)

type OrderLine struct {
	ProductID int64
	Quantity  int
}

// Order: sus líneas se cargan al primer acceso (lazy initialization)
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64

	db        *sql.DB
	linesOnce sync.Once
	lines     []OrderLine
	linesErr  error
}

// Finder: carga el pedido SIN sus líneas
func FindOrder(ctx context.Context, db *sql.DB, id int64) (*Order, error) {
	o := &Order{db: db}
	err := db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents
		 FROM orders WHERE id = $1`, id).
		Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Lines: Lazy Load — la consulta a order_lines ocurre UNA vez,
// en el primer acceso
func (o *Order) Lines(ctx context.Context) ([]OrderLine, error) {
	o.linesOnce.Do(func() {
		rows, err := o.db.QueryContext(ctx,
			`SELECT product_id, quantity
			 FROM order_lines WHERE order_id = $1`, o.ID)
		if err != nil {
			o.linesErr = err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var l OrderLine
			if err := rows.Scan(&l.ProductID, &l.Quantity); err != nil {
				o.linesErr = err
				return
			}
			o.lines = append(o.lines, l)
		}
		o.linesErr = rows.Err()
	})
	return o.lines, o.linesErr
}
