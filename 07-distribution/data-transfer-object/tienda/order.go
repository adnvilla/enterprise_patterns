package tienda

import (
	"errors"
	"time"
)

// Modelo de dominio: rico, con campos que son asunto interno
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	Version    int64 // optimistic locking: a nadie le importa afuera
	CostCents  int64 // costo interno: ¡esto NO se publica!
	CreatedAt  time.Time
	Lines      []OrderLine
}

type OrderLine struct {
	ProductID int64
	Name      string
	Quantity  int
	UnitCents int64
}

func (o *Order) TotalCents() int64 {
	var total int64
	for _, l := range o.Lines {
		total += int64(l.Quantity) * l.UnitCents
	}
	return total
}

var ErrSinLineas = errors.New("un pedido necesita al menos una línea")
