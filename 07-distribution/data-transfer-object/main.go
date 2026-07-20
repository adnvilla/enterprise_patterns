package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/07-distribution/data-transfer-object/tienda"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		// Entrada: JSON -> request DTO -> dominio
		var req tienda.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		order, err := req.ToDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		order.ID = 42 // aquí iría la persistencia real

		// Salida: dominio -> DTO -> JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tienda.AssembleOrderDTO(order))
	})

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Ruta de ejemplo:")
	fmt.Println(`  curl -X POST http://localhost:8080/orders -H 'Content-Type: application/json' \`)
	fmt.Println(`    -d '{"customer_id":7,"lines":[{"product_id":1,"quantity":2}]}'`)
	fmt.Println()

	// Demostración: el cliente manda un request DTO y recibe un OrderDTO.
	// Version y CostCents jamás aparecen en la respuesta.
	body, _ := json.Marshal(tienda.CreateOrderRequest{
		CustomerID: 7,
		Lines: []tienda.CreateOrderLineRequest{
			{ProductID: 1, Quantity: 2},
			{ProductID: 5, Quantity: 1},
		},
	})
	resp, err := http.Post("http://localhost:8080/orders", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, _ := io.ReadAll(resp.Body)
	fmt.Printf("Respuesta de la API (solo el DTO, sin Version ni CostCents):\n%s", out)
}
