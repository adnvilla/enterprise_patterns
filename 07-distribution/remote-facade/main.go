package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/07-distribution/remote-facade/tienda"
)

// ANTI-EJEMPLO: exponer la interfaz fina por red (no hagas esto).
// Cada método del dominio convertido en endpoint:
//
//   POST /orders            -> crea el pedido vacío      (viaje 1)
//   POST /orders/42/lines   -> agrega la línea 1         (viaje 2)
//   POST /orders/42/lines   -> agrega la línea 2         (viaje 3)
//   POST /orders/42/lines   -> agrega la línea 3         (viaje 4)
//   POST /orders/42/pay     -> cobra                     (viaje 5)
//
// El cliente termina haciendo esto:
//
//	for _, linea := range lineas {
//		// un viaje de red completo POR CADA línea del pedido
//		http.Post(base+"/orders/42/lines", "application/json", codifica(linea))
//	}
//
// Con 30 líneas: 32 viajes, 32 latencias, 32 oportunidades
// de fallar dejando un pedido a medias. La fachada gruesa
// hace el mismo trabajo en UN viaje.

func main() {
	// Servidor: expone la fachada gruesa
	http.HandleFunc("/orders/place", tienda.PlaceOrderFacade)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Ruta de ejemplo:")
	fmt.Println(`  curl -X POST http://localhost:8080/orders/place -H 'Content-Type: application/json' \`)
	fmt.Println(`    -d '{"customer_id":7,"lines":[{"product_id":1,"quantity":2,"unit_cents":12900}],"payment_cents":25800}'`)
	fmt.Println()

	// Cliente remoto: UNA sola petición con todo el pedido
	pedido := map[string]any{
		"customer_id": 7,
		"lines": []map[string]any{
			{"product_id": 1, "quantity": 2, "unit_cents": 12900},
			{"product_id": 5, "quantity": 1, "unit_cents": 45900},
		},
		"payment_cents": 71700,
	}
	body, _ := json.Marshal(pedido)

	resp, err := http.Post("http://localhost:8080/orders/place", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var out tienda.PlaceOrderResponse
	json.NewDecoder(resp.Body).Decode(&out)
	fmt.Printf("Pedido %d: %s, total %d centavos\n", out.OrderID, out.Status, out.TotalCents)
}
