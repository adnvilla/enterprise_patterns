package pedidos

import "database/sql"

// Row Data Gateway: cada instancia ES una fila de orders.
type OrderRowGateway struct {
	db         *sql.DB
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

// Finder: busca la fila y construye su gateway.
func FindOrderRow(db *sql.DB, id int64) (*OrderRowGateway, error) {
	r := &OrderRowGateway{db: db}
	err := db.QueryRow(
		`SELECT id, customer_id, status, total_cents
		   FROM orders WHERE id = $1`, id,
	).Scan(&r.ID, &r.CustomerID, &r.Status, &r.TotalCents)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Update persiste los cambios de ESTA fila, no de otra.
func (r *OrderRowGateway) Update() error {
	_, err := r.db.Exec(
		`UPDATE orders SET status = $2, total_cents = $3 WHERE id = $1`,
		r.ID, r.Status, r.TotalCents)
	return err
}

// Delete elimina ESTA fila.
func (r *OrderRowGateway) Delete() error {
	_, err := r.db.Exec(`DELETE FROM orders WHERE id = $1`, r.ID)
	return err
}
