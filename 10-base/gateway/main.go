package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/adnvilla/enterprise_patterns/10-base/gateway/pagos"
)

// API externa SIMULADA: hace las veces del proveedor de pagos real.
// Habla el formato PROPIO del proveedor ("amount", "status", "tx_id"),
// que jamás sale del paquete pagos.
func proveedorSimulado() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/charges", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Amount int64  `json:"amount"`
			Ref    string `json:"external_reference"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status := "ok"
		if req.Amount > 500000 { // el proveedor rechaza cobros enormes
			status = "declined"
		}
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": status,
			"tx_id":  "tx-" + req.Ref,
		})
	})
	return httptest.NewServer(mux)
}

func main() {
	srv := proveedorSimulado()
	defer srv.Close()

	gw := pagos.NewPaymentGateway(srv.URL, "sk_test_123")

	// El cliente invoca al Gateway CONSCIENTEMENTE, en términos del dominio;
	// el Mapper trabajó por dentro sin que nadie lo viera
	charge, err := gw.Charge(context.Background(), 42, 129900)
	if err != nil {
		log.Fatal(err)
	}

	if charge.Approved {
		fmt.Println("Pedido cobrado, referencia:", charge.Reference)
	} else {
		fmt.Println("Pago rechazado para el pedido", charge.OrderID)
	}

	// Un cobro enorme, para ver al proveedor rechazando
	charge, err = gw.Charge(context.Background(), 43, 999999)
	if err != nil {
		log.Fatal(err)
	}

	if charge.Approved {
		fmt.Println("Pedido cobrado, referencia:", charge.Reference)
	} else {
		fmt.Println("Pago rechazado para el pedido", charge.OrderID)
	}
}
