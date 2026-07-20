package orders

import "database/sql"

// Table Module: UNA instancia que responde por TODAS las filas
// de la tabla orders. No hay un objeto por pedido.
type OrdersModule struct {
	db *sql.DB
}

func NewOrdersModule(db *sql.DB) *OrdersModule {
	return &OrdersModule{db: db}
}

// TotalByCustomer opera sobre el conjunto: agrega todas las filas
// del cliente con SQL, sin materializar objetos pedido.
func (m *OrdersModule) TotalByCustomer(customerID int64) (int64, error) {
	var total sql.NullInt64
	err := m.db.QueryRow(
		`SELECT SUM(total_cents)
		   FROM orders
		  WHERE customer_id = $1
		    AND status <> 'cancelado'`,
		customerID,
	).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total.Int64, nil // NULL (sin pedidos) se convierte en 0
}

// MarkOverdue es lógica de negocio aplicada en bloque:
// todo pedido pendiente con más de `days` días se marca vencido.
func (m *OrdersModule) MarkOverdue(days int) (int64, error) {
	res, err := m.db.Exec(
		`UPDATE orders
		    SET status = 'vencido'
		  WHERE status = 'pendiente'
		    AND created_at < now() - make_interval(days => $1)`,
		days,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected() // cuántas filas procesó el módulo
}
