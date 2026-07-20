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

// Unit of Work: registra nuevos, sucios y eliminados
type UnitOfWork struct {
	db            *sql.DB
	newOrders     []*Order
	dirtyOrders   []*Order
	removedOrders []*Order
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

// RegisterNew: el objeto se creó en esta transacción de negocio
func (u *UnitOfWork) RegisterNew(o *Order) { u.newOrders = append(u.newOrders, o) }

// RegisterDirty: el objeto existe pero fue modificado
func (u *UnitOfWork) RegisterDirty(o *Order) { u.dirtyOrders = append(u.dirtyOrders, o) }

// RegisterRemoved: el objeto debe eliminarse
func (u *UnitOfWork) RegisterRemoved(o *Order) { u.removedOrders = append(u.removedOrders, o) }

// Commit: UNA transacción para todos los cambios acumulados
func (u *UnitOfWork) Commit(ctx context.Context) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // si algo falla, no se escribe nada

	for _, o := range u.newOrders {
		err = tx.QueryRowContext(ctx,
			`INSERT INTO orders (customer_id, status, total_cents, version)
			 VALUES ($1, $2, $3, 1) RETURNING id`,
			o.CustomerID, o.Status, o.TotalCents).Scan(&o.ID)
		if err != nil {
			return err
		}
	}
	for _, o := range u.dirtyOrders {
		_, err = tx.ExecContext(ctx,
			`UPDATE orders SET status = $1, total_cents = $2, version = version + 1
			 WHERE id = $3`,
			o.Status, o.TotalCents, o.ID)
		if err != nil {
			return err
		}
	}
	for _, o := range u.removedOrders {
		_, err = tx.ExecContext(ctx, `DELETE FROM orders WHERE id = $1`, o.ID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
