package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/01-domain-logic/service-layer/app"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://app:app@localhost:5432/tienda?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := app.NewOrderService(db)

	// El handler HTTP no sabe de reglas ni de transacciones:
	// traduce la petición y delega en el Service Layer.
	http.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			CustomerID int64                  `json:"customer_id"`
			Lines      []app.OrderLineRequest `json:"lines"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		orderID, err := service.PlaceOrder(r.Context(), req.CustomerID, req.Lines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"order_id": orderID})
	})

	// Levantamos el servidor y hacemos una petición de demostración
	// end-to-end contra la base de datos del compose.
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := http.Serve(ln, nil); err != nil {
			log.Println(err)
		}
	}()

	resp, err := http.Post("http://127.0.0.1:8080/orders", "application/json",
		strings.NewReader(`{
			"customer_id": 42,
			"lines": [
				{"product_id": 1, "quantity": 2},
				{"product_id": 3, "quantity": 1}
			]
		}`))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("POST /orders -> %s: %s", resp.Status, body)
}
