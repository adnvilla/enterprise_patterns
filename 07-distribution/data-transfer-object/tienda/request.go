package tienda

import "time"

// Request DTO: lo que el cliente PUEDE mandar, y nada más.
// Nadie puede colarte un "status": "paid" ni un CostCents,
// porque el struct simplemente no tiene dónde ponerlos.
type CreateOrderRequest struct {
	CustomerID int64                    `json:"customer_id"`
	Lines      []CreateOrderLineRequest `json:"lines"`
}

type CreateOrderLineRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

// Del DTO de entrada al dominio, validando en la frontera
func (r CreateOrderRequest) ToDomain() (*Order, error) {
	if len(r.Lines) == 0 {
		return nil, ErrSinLineas
	}
	o := &Order{
		CustomerID: r.CustomerID,
		Status:     "pending", // el estado inicial lo decide el dominio, no el cliente
		CreatedAt:  time.Now(),
	}
	for _, l := range r.Lines {
		o.Lines = append(o.Lines, OrderLine{ProductID: l.ProductID, Quantity: l.Quantity})
	}
	return o, nil
}
