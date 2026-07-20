package apiview

import (
	"encoding/json"
	"net/http"
	"time"
)

func OrdersAPIHandler(w http.ResponseWriter, r *http.Request) {
	orders := []Order{ // normalmente vendrían del repositorio
		{ID: 1, Status: "PAID", TotalCents: 129900, CreatedAt: time.Now()},
		{ID: 2, Status: "PENDING", TotalCents: 45000, CreatedAt: time.Now()},
	}

	// La vista ES la transformación: elemento por elemento.
	out := make([]OrderJSON, 0, len(orders))
	for _, o := range orders {
		out = append(out, ToOrderJSON(o))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}
