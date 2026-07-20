package pedidos

import "database/sql"

// OrderRow es solo un registro plano: sin comportamiento.
type OrderRow struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

// Table Data Gateway: UNA instancia atiende TODA la tabla orders.
type OrdersGateway struct {
	db *sql.DB
}

func NewOrdersGateway(db *sql.DB) *OrdersGateway {
	return &OrdersGateway{db: db}
}

// FindByCustomer devuelve datos planos, no objetos de dominio.
func (g *OrdersGateway) FindByCustomer(customerID int64) ([]OrderRow, error) {
	rows, err := g.db.Query(
		`SELECT id, customer_id, status, total_cents
		   FROM orders WHERE customer_id = $1`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []OrderRow
	for rows.Next() {
		var r OrderRow
		if err := rows.Scan(&r.ID, &r.CustomerID, &r.Status, &r.TotalCents); err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, rows.Err()
}

// Insert recibe valores sueltos y regresa el id generado.
func (g *OrdersGateway) Insert(customerID int64, status string, totalCents int64) (int64, error) {
	var id int64
	err := g.db.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1) RETURNING id`,
		customerID, status, totalCents).Scan(&id)
	return id, err
}
