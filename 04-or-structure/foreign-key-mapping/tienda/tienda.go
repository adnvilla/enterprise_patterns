package tienda

import "database/sql"

type Order struct {
	ID         int64
	CustomerID int64 // la FK, tal cual, en el objeto hijo
	Status     string
	TotalCents int64
}

type Customer struct {
	ID     int64
	Name   string
	Email  string
	Orders []*Order // en memoria, el padre tiene la colección
}

type CustomerMapper struct {
	DB *sql.DB
}

// FindWithOrders resuelve la dirección fila -> objeto:
// una consulta por el padre y otra por "las filas cuya FK soy yo".
func (m *CustomerMapper) FindWithOrders(id int64) (*Customer, error) {
	c := &Customer{}
	err := m.DB.QueryRow(
		`SELECT id, name, email FROM customers WHERE id = $1`, id,
	).Scan(&c.ID, &c.Name, &c.Email)
	if err != nil {
		return nil, err
	}

	rows, err := m.DB.Query(
		`SELECT id, customer_id, status, total_cents
		 FROM orders WHERE customer_id = $1
		 ORDER BY id`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		o := &Order{}
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents); err != nil {
			return nil, err
		}
		c.Orders = append(c.Orders, o)
	}
	return c, rows.Err()
}

// La dirección objeto -> fila es simplemente escribir el Identity Field
// del padre en la columna FK del hijo al insertarlo.
func (m *CustomerMapper) AddOrder(c *Customer, o *Order) error {
	o.CustomerID = c.ID // el hijo apunta al padre
	err := m.DB.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1) RETURNING id`,
		o.CustomerID, o.Status, o.TotalCents,
	).Scan(&o.ID)
	if err == nil {
		c.Orders = append(c.Orders, o)
	}
	return err
}
