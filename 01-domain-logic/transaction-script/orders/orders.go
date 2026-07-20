package orders

import (
	"database/sql"
	"errors"
)

type Line struct {
	ProductID int64
	Quantity  int
}

// Transaction Script: el caso de uso «colocar pedido» completo,
// de la validación al INSERT, en un solo procedimiento.
func PlaceOrder(db *sql.DB, customerID int64, lines []Line) (int64, error) {
	if len(lines) == 0 {
		return 0, errors.New("el pedido necesita al menos una línea")
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // no hace nada si ya hubo Commit

	// 1) Validar y calcular el total con los precios reales.
	var total int64
	for _, l := range lines {
		if l.Quantity <= 0 {
			return 0, errors.New("la cantidad debe ser mayor que cero")
		}
		var price int64
		err := tx.QueryRow(
			`SELECT price_cents FROM products WHERE id = $1`, l.ProductID,
		).Scan(&price)
		if err != nil {
			return 0, err
		}
		total += price * int64(l.Quantity)
	}

	// 2) Insertar el pedido.
	var orderID int64
	err = tx.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, 'nuevo', $2, 1)
		 RETURNING id`,
		customerID, total,
	).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	// 3) Insertar las líneas del pedido.
	for _, l := range lines {
		_, err = tx.Exec(
			`INSERT INTO order_lines (order_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			orderID, l.ProductID, l.Quantity,
		)
		if err != nil {
			return 0, err
		}
	}

	return orderID, tx.Commit()
}
