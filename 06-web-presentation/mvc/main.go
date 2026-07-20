package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/mvc/orders"
)

func main() {
	repo := orders.NewRepository()

	mux := http.NewServeMux()
	// El mux dirige la ruta a su controlador.
	mux.Handle("GET /orders", &orders.ListController{Repo: repo})

	log.Println("Escuchando en :8080. Prueba con:")
	log.Println("  curl http://localhost:8080/orders")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
