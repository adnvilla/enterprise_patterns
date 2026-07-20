package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/page-controller/orders"
)

func main() {
	repo := orders.NewRepository()

	// El mux estándar de Go ES el mapeo página → Page Controller.
	mux := http.NewServeMux()
	mux.Handle("GET /orders", &orders.OrdersHandler{Repo: repo})
	mux.Handle("GET /orders/detail", &orders.OrderDetailHandler{Repo: repo})

	log.Println("Escuchando en :8080. Prueba con:")
	log.Println("  curl http://localhost:8080/orders")
	log.Println("  curl http://localhost:8080/orders/detail?id=1")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
