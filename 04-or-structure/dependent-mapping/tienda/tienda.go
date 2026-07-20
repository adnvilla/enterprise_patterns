package tienda

import "database/sql"

// OrderLine no tiene id ni mapper propio: es un dependiente de Order.
type OrderLine struct {
	ProductID int64
	Quantity  int
}

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
	Lines      []OrderLine
}

type OrderMapper struct {
	DB *sql.DB
}

// Update persiste al dueño Y a sus dependientes en una transacción.
// Con dependientes, lo simple gana: borrar e insertar las líneas.
func (m *OrderMapper) Update(o *Order) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(
		`UPDATE orders SET status = $1, total_cents = $2 WHERE id = $3`,
		o.Status, o.TotalCents, o.ID); err != nil {
		return err
	}

	// El mapper del padre es el único que toca order_lines.
	if _, err := tx.Exec(
		`DELETE FROM order_lines WHERE order_id = $1`, o.ID); err != nil {
		return err
	}
	for _, l := range o.Lines {
		if _, err := tx.Exec(
			`INSERT INTO order_lines (order_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			o.ID, l.ProductID, l.Quantity); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// FindByID carga al dueño CON sus dependientes: nadie lee una línea suelta.
func (m *OrderMapper) FindByID(id int64) (*Order, error) {
	o := &Order{}
	err := m.DB.QueryRow(
		`SELECT id, customer_id, status, total_cents
		 FROM orders WHERE id = $1`, id,
	).Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}

	rows, err := m.DB.Query(
		`SELECT product_id, quantity
		 FROM order_lines WHERE order_id = $1`, id)
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
