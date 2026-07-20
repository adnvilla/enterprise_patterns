package memoria

import (
	"context"
	"fmt"

	"github.com/adnvilla/enterprise_patterns/10-base/separated-interface/dominio" // la flecha de dependencia apunta AL dominio
)

// Implementación en memoria de dominio.OrderRepository: el mismo contrato
// que la de postgres, sin base de datos — ideal para demos y pruebas
type OrderRepository struct {
	orders map[int64]dominio.Order
	nextID int64
}

// También satisface la interfaz del dominio implícitamente
var _ dominio.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{orders: make(map[int64]dominio.Order)}
}

func (r *OrderRepository) ByID(ctx context.Context, id int64) (*dominio.Order, error) {
	o, ok := r.orders[id]
	if !ok {
		return nil, fmt.Errorf("pedido %d no encontrado", id)
	}
	return &o, nil
}

func (r *OrderRepository) Save(ctx context.Context, o *dominio.Order) error {
	if o.IsNew() { // cortesía del Layer Supertype
		r.nextID++
		o.ID = r.nextID
	}
	r.orders[o.ID] = *o
	return nil
}
