package tienda

import (
	"encoding/json"
	"net/http"
)

// Lo que entra y sale de la fachada: datos planos, todo junto
type PlaceOrderRequest struct {
	CustomerID int64 `json:"customer_id"`
	Lines      []struct {
		ProductID int64 `json:"product_id"`
		Quantity  int   `json:"quantity"`
		UnitCents int64 `json:"unit_cents"`
	} `json:"lines"`
	PaymentCents int64 `json:"payment_cents"`
}

type PlaceOrderResponse struct {
	OrderID    int64  `json:"order_id"`
	Status     string `json:"status"`
	TotalCents int64  `json:"total_cents"`
}

// Remote Facade: UNA petición remota, muchas llamadas locales.
// Nota: aquí no hay reglas de negocio, solo traducción y coordinación.
func PlaceOrderFacade(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "petición inválida", http.StatusBadRequest)
		return
	}

	// Por dentro, la fachada coordina el modelo fino
	order := &Order{CustomerID: req.CustomerID, Status: "pending"}
	for _, l := range req.Lines {
		order.AddLine(l.ProductID, l.Quantity, l.UnitCents) // llamada local: gratis
	}
	if err := order.Pay(req.PaymentCents); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	order.ID = 42 // aquí iría la persistencia real (repositorio, Unit of Work...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PlaceOrderResponse{
		OrderID: order.ID, Status: order.Status, TotalCents: order.TotalCents(),
	})
}
