package tienda

// DTO: solo datos, solo lo que el cliente debe ver.
// Los tags json definen el contrato público de la API.
type OrderDTO struct {
	ID         int64          `json:"id"`
	Status     string         `json:"status"`
	TotalCents int64          `json:"total_cents"`
	CreatedAt  string         `json:"created_at"`
	Lines      []OrderLineDTO `json:"lines"`
}

type OrderLineDTO struct {
	ProductID int64  `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitCents int64  `json:"unit_cents"`
}
