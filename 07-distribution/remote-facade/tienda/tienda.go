package tienda

import "errors"

// Modelo fino: objetos pequeños, métodos pequeños.
// Perfectos en memoria; letales si los expones por red uno a uno.
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	lines      []OrderLine
}

type OrderLine struct {
	ProductID int64
	Quantity  int
	UnitCents int64
}

// Métodos finos: cada llamada local cuesta nanosegundos
func (o *Order) AddLine(productID int64, qty int, unitCents int64) {
	o.lines = append(o.lines, OrderLine{ProductID: productID, Quantity: qty, UnitCents: unitCents})
}

func (o *Order) TotalCents() int64 {
	var total int64
	for _, l := range o.lines {
		total += int64(l.Quantity) * l.UnitCents
	}
	return total
}

func (o *Order) Pay(amountCents int64) error {
	if amountCents < o.TotalCents() {
		return errors.New("pago insuficiente")
	}
	o.Status = "paid"
	return nil
}
