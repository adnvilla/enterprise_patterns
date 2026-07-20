package tienda

import "time"

// Assembler: la única pieza que conoce ambos mundos
func AssembleOrderDTO(o *Order) OrderDTO {
	lines := make([]OrderLineDTO, 0, len(o.Lines))
	for _, l := range o.Lines {
		lines = append(lines, OrderLineDTO{
			ProductID: l.ProductID,
			Name:      l.Name,
			Quantity:  l.Quantity,
			UnitCents: l.UnitCents,
		})
	}
	return OrderDTO{
		ID:         o.ID,
		Status:     o.Status,
		TotalCents: o.TotalCents(), // dato calculado: el cliente no suma
		CreatedAt:  o.CreatedAt.Format(time.RFC3339),
		Lines:      lines,
	}
	// Fíjate en lo que NO está: Version y CostCents nunca salen de aquí
}
