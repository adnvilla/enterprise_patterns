package dominio

import "errors"

// Objeto de dominio: puro negocio, ignorante de la base de datos.
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

func NewOrder(customerID int64, totalCents int64) *Order {
	return &Order{CustomerID: customerID, Status: "pending", TotalCents: totalCents}
}

// Las reglas de negocio viven aquí, y se prueban sin base de datos.
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

func (o *Order) ApplyDiscount(percent int64) error {
	if percent <= 0 || percent > 50 {
		return errors.New("el descuento debe estar entre 1% y 50%")
	}
	o.TotalCents -= o.TotalCents * percent / 100
	return nil
}
