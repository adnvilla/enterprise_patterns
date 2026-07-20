package pedidos

import (
	"database/sql"
	"errors"
)

// Active Record: la fila de orders, con comportamiento y persistencia.
type Order struct {
	db         *sql.DB
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

func NewOrder(db *sql.DB, customerID int64, totalCents int64) *Order {
	return &Order{db: db, CustomerID: customerID, Status: "pending", TotalCents: totalCents}
}

// Lógica de negocio simple: vive DENTRO del propio objeto.
func (o *Order) MarkAsPaid() error {
	if o.Status != "pending" {
		return errors.New("solo un pedido pendiente puede pagarse")
	}
	if o.TotalCents <= 0 {
		return errors.New("un pedido sin importe no puede pagarse")
	}
	o.Status = "paid"
	return nil
}

// Insert: el objeto se guarda a sí mismo.
func (o *Order) Insert() error {
	return o.db.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1) RETURNING id`,
		o.CustomerID, o.Status, o.TotalCents).Scan(&o.ID)
}

// Update: el objeto actualiza SU fila.
func (o *Order) Update() error {
	_, err := o.db.Exec(
		`UPDATE orders SET status = $2, total_cents = $3 WHERE id = $1`,
		o.ID, o.Status, o.TotalCents)
	return err
}

// Delete: el objeto borra SU fila.
func (o *Order) Delete() error {
	_, err := o.db.Exec(`DELETE FROM orders WHERE id = $1`, o.ID)
	return err
}

// Finder: hidrata un Active Record desde la base.
func FindOrder(db *sql.DB, id int64) (*Order, error) {
	o := &Order{db: db}
	err := db.QueryRow(
		`SELECT id, customer_id, status, total_cents
		   FROM orders WHERE id = $1`, id,
	).Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents)
	if err != nil {
		return nil, err
	}
	return o, nil
}
