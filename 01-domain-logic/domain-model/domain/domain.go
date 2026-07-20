package domain

import "errors"

// Value Object: una línea se define por sus valores, no por una identidad.
type OrderLine struct {
	ProductID  int64
	Quantity   int
	PriceCents int64 // precio unitario congelado al agregar la línea
}

// Entity: el pedido tiene identidad y ciclo de vida propios.
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	lines      []OrderLine
	discount   int64 // descuento acumulado en centavos
}

func NewOrder(customerID int64) *Order {
	return &Order{CustomerID: customerID, Status: "nuevo"}
}

// AddLine protege las invariantes del negocio:
// nada de líneas inválidas ni pedidos confirmados que cambian.
func (o *Order) AddLine(productID int64, quantity int, priceCents int64) error {
	if o.Status != "nuevo" {
		return errors.New("no se pueden agregar líneas a un pedido confirmado")
	}
	if quantity <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}
	o.lines = append(o.lines, OrderLine{
		ProductID:  productID,
		Quantity:   quantity,
		PriceCents: priceCents,
	})
	return nil
}

func (o *Order) subtotal() int64 {
	var total int64
	for _, l := range o.lines {
		total += l.PriceCents * int64(l.Quantity)
	}
	return total
}

// Total es una regla de negocio, no una consulta SQL.
func (o *Order) Total() int64 {
	total := o.subtotal() - o.discount
	if total < 0 {
		total = 0
	}
	return total
}

// ApplyDiscount: 10% para pedidos de $100.00 (10000 centavos) o más.
// La regla vive aquí, y solo aquí.
func (o *Order) ApplyDiscount() {
	if o.subtotal() >= 10000 {
		o.discount = o.subtotal() / 10
	}
}

func (o *Order) Confirm() error {
	if len(o.lines) == 0 {
		return errors.New("un pedido vacío no se puede confirmar")
	}
	o.Status = "confirmado"
	return nil
}

// Lines expone una copia para que nadie mutile el estado interno.
func (o *Order) Lines() []OrderLine {
	out := make([]OrderLine, len(o.lines))
	copy(out, o.lines)
	return out
}
