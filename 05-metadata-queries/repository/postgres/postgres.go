package postgres

import (
	"database/sql"

	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/domain"
)

// Concrete Repository: aquí — y SOLO aquí — vive el SQL.
type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) Add(o *domain.Order) error {
	return r.DB.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1)
		 RETURNING id`,
		o.CustomerID, o.Status, o.TotalCents,
	).Scan(&o.ID)
}

func (r *OrderRepository) PaidByCustomer(customerID int64) ([]domain.Order, error) {
	rows, err := r.DB.Query(
		`SELECT id, customer_id, status, total_cents
		   FROM orders
		  WHERE customer_id = $1 AND status = 'paid'`,
		customerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}
