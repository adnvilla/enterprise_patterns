package app

import (
	"context"
	"database/sql"

	"github.com/adnvilla/enterprise_patterns/01-domain-logic/service-layer/domain"
)

// Service Layer: fachada de casos de uso sobre el dominio.
type OrderService struct {
	db *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{db: db}
}

type OrderLineRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

// PlaceOrder es lógica de APLICACIÓN: define el límite transaccional
// y coordina dominio + infraestructura. Las reglas de NEGOCIO
// (descuento, invariantes) viven en el dominio, no aquí.
func (s *OrderService) PlaceOrder(ctx context.Context, customerID int64, lines []OrderLineRequest) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // no hace nada si ya hubo Commit

	// 1) Armar el pedido con el Domain Model de la entrada anterior.
	order := domain.NewOrder(customerID)
	for _, l := range lines {
		var price int64
		if err := tx.QueryRowContext(ctx,
			`SELECT price_cents FROM products WHERE id = $1`, l.ProductID,
		).Scan(&price); err != nil {
			return 0, err
		}
		// El dominio valida cantidades y estado del pedido.
		if err := order.AddLine(l.ProductID, l.Quantity, price); err != nil {
			return 0, err
		}
	}
	order.ApplyDiscount() // regla de negocio: decide el dominio
	if err := order.Confirm(); err != nil {
		return 0, err
	}

	// 2) Persistir el resultado (más adelante en la serie,
	//    esto será trabajo de un Data Mapper / Repository).
	var orderID int64
	if err := tx.QueryRowContext(ctx,
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1)
		 RETURNING id`,
		order.CustomerID, order.Status, order.Total(),
	).Scan(&orderID); err != nil {
		return 0, err
	}
	for _, l := range order.Lines() {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO order_lines (order_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			orderID, l.ProductID, l.Quantity,
		); err != nil {
			return 0, err
		}
	}

	// 3) El servicio decide dónde termina la transacción.
	return orderID, tx.Commit()
}
