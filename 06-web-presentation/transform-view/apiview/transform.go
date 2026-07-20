package apiview

import "time"

// Modelo de dominio: la verdad interna de la aplicación.
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
	Version    int64
	CreatedAt  time.Time
}

// DTO de salida: la forma PÚBLICA de la respuesta. Nota que
// no expone Version ni CustomerID: eso es asunto interno.
type OrderJSON struct {
	ID     int64   `json:"id"`
	Status string  `json:"status"`
	Total  float64 `json:"total"`
	Date   string  `json:"date"`
}

// Transformer: convierte cada elemento del modelo en su salida.
func ToOrderJSON(o Order) OrderJSON {
	return OrderJSON{
		ID:     o.ID,
		Status: o.Status,
		Total:  float64(o.TotalCents) / 100,
		Date:   o.CreatedAt.Format("2006-01-02"),
	}
}
