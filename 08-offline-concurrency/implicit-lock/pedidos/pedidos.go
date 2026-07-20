package pedidos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrConflict = errors.New("conflicto de concurrencia: el pedido fue modificado por alguien más")

// El objeto de dominio NO expone versión ni lock alguno:
// el control de concurrencia es asunto de la capa de datos, no del negocio
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

// Repository con Implicit Lock: actúa como el Unit of Work de una
// transacción de negocio. Al leer registra la versión, al guardar la
// verifica — SIEMPRE, sin que nadie tenga que acordarse de pedirlo.
type OrderRepository struct {
	db       *sql.DB
	versions map[int64]int64 // versión leída de cada pedido en ESTA sesión
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db, versions: make(map[int64]int64)}
}

// FindByID: lee el pedido y adquiere el lock optimista automáticamente
// (registra por dentro la versión leída; el llamador nunca la ve)
func (r *OrderRepository) FindByID(ctx context.Context, id int64) (*Order, error) {
	o := &Order{}
	var version int64
	err := r.db.QueryRowContext(ctx,
		`SELECT id, customer_id, status, total_cents, version
		 FROM orders WHERE id = $1`, id).
		Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents, &version)
	if err != nil {
		return nil, err
	}
	// Implicit Lock: la adquisición ocurre aquí, no en el código de negocio
	r.versions[id] = version
	return o, nil
}

// Save: única puerta de escritura del pedido. El chequeo de versión vive
// DENTRO, así que olvidarlo es imposible por construcción.
func (r *OrderRepository) Save(ctx context.Context, o *Order) error {
	version, ok := r.versions[o.ID]
	if !ok {
		// Ni siquiera se puede guardar algo que no pasó por FindByID:
		// no existe una ruta de escritura que esquive el lock
		return fmt.Errorf("el pedido %d no fue leído por este repositorio: no hay lock que verificar", o.ID)
	}
	res, err := r.db.ExecContext(ctx,
		`UPDATE orders
		 SET status = $1, total_cents = $2, version = version + 1
		 WHERE id = $3 AND version = $4`,
		o.Status, o.TotalCents, o.ID, version)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// 0 filas afectadas = alguien se nos adelantó
	if affected == 0 {
		return ErrConflict
	}
	r.versions[o.ID] = version + 1 // la sesión sigue al día
	return nil
}
