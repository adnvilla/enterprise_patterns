package persistencia

import (
	"database/sql"

	"github.com/adnvilla/enterprise_patterns/02-data-source/data-mapper/dominio"
)

// Data Mapper: el ÚNICO que conoce el SQL y el dominio a la vez.
type OrderMapper struct {
	db *sql.DB
}

func NewOrderMapper(db *sql.DB) *OrderMapper {
	return &OrderMapper{db: db}
}

// Find hidrata un objeto de dominio desde su fila.
func (m *OrderMapper) Find(id int64) (*dominio.Order, error) {
	o := &dominio.Order{}
	err := m.db.QueryRow(
		`SELECT id, customer_id, status, total_cents
		   FROM orders WHERE id = $1`, id,
	).Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Insert traduce el objeto a una fila nueva y le asigna su id.
func (m *OrderMapper) Insert(o *dominio.Order) error {
	return m.db.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1) RETURNING id`,
		o.CustomerID, o.Status, o.TotalCents).Scan(&o.ID)
}

// Update vuelca el estado actual del objeto sobre su fila.
func (m *OrderMapper) Update(o *dominio.Order) error {
	_, err := m.db.Exec(
		`UPDATE orders SET status = $2, total_cents = $3 WHERE id = $1`,
		o.ID, o.Status, o.TotalCents)
	return err
}
