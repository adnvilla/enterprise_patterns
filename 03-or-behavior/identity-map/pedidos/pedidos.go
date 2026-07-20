package pedidos

import (
	"context"
	"database/sql"
)

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

// Finder con Identity Map: una sola instancia por id dentro de la sesión
type OrderFinder struct {
	db     *sql.DB
	loaded map[int64]*Order // Identity Map: id -> objeto
}

func NewOrderFinder(db *sql.DB) *OrderFinder {
	return &OrderFinder{db: db, loaded: make(map[int64]*Order)}
}

func (f *OrderFinder) Find(ctx context.Context, id int64) (*Order, error) {
	// 1) Primero el mapa: si ya lo cargamos, devolvemos LA MISMA instancia
	if o, ok := f.loaded[id]; ok {
		return o, nil
	}

	// 2) Si no está, vamos a la base de datos...
	o := &Order{}
	err := f.db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents
		 FROM orders WHERE id = $1`, id).
		Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}

	// 3) ...y lo registramos antes de devolverlo
	f.loaded[id] = o
	return o, nil
}
