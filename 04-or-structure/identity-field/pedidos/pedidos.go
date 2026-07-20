package pedidos

import "database/sql"

// Order guarda el id de su fila: ese campo es el Identity Field.
type Order struct {
	ID         int64 // 0 = "aún no persistido"
	CustomerID int64
	Status     string
	TotalCents int64
}

// IsNew aprovecha el valor cero de Go: un id 0 nunca sale de BIGSERIAL,
// así que sirve como bandera de "este objeto todavía no tiene fila".
func (o *Order) IsNew() bool { return o.ID == 0 }

// OrderMapper conoce la correspondencia objeto <-> fila.
type OrderMapper struct {
	DB *sql.DB
}

// Insert persiste el pedido y puebla el Identity Field con RETURNING:
// una sola ida a la base para insertar y conocer el id.
func (m *OrderMapper) Insert(o *Order) error {
	return m.DB.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1)
		 RETURNING id`,
		o.CustomerID, o.Status, o.TotalCents,
	).Scan(&o.ID) // el id generado aterriza directo en el struct
}

// FindByID usa el Identity Field para localizar la fila exacta.
func (m *OrderMapper) FindByID(id int64) (*Order, error) {
	o := &Order{}
	err := m.DB.QueryRow(
		`SELECT id, customer_id, status, total_cents
		 FROM orders WHERE id = $1`,
		id,
	).Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}
	return o, nil
}
