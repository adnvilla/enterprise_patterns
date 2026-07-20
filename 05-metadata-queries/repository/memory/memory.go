package memory

import "github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/domain"

// In-Memory Repository: la misma interfaz, cero base de datos.
// Para el dominio es indistinguible de la versión con Postgres.
type OrderRepository struct {
	orders []domain.Order
	nextID int64
}

func (r *OrderRepository) Add(o *domain.Order) error {
	r.nextID++
	o.ID = r.nextID
	r.orders = append(r.orders, *o)
	return nil
}

func (r *OrderRepository) PaidByCustomer(customerID int64) ([]domain.Order, error) {
	var result []domain.Order
	for _, o := range r.orders {
		if o.CustomerID == customerID && o.Status == "paid" {
			result = append(result, o)
		}
	}
	return result, nil
}
