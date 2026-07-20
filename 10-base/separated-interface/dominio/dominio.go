package dominio

import "context"

// Layer Supertype: el comportamiento común de toda la capa de dominio
type Entity struct {
	ID int64
}

// IsNew: aún no ha sido persistido
func (e Entity) IsNew() bool { return e.ID == 0 }

// En Go no heredamos: embebemos
type Order struct {
	Entity // Order "es un" objeto de dominio con ID
	CustomerID int64
	Status     string
	TotalCents int64
}

// Separated Interface: el CONSUMIDOR define el contrato.
// El dominio no importa la infraestructura; la infraestructura importará al dominio.
type OrderRepository interface {
	ByID(ctx context.Context, id int64) (*Order, error)
	Save(ctx context.Context, o *Order) error
}
