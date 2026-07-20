package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/transform-view/apiview"
)

func main() {
	// Transform View: la API JSON.
	http.HandleFunc("/api/ordenes", apiview.OrdersAPIHandler)

	log.Println("Escuchando en :8080. Prueba con:")
	log.Println("  curl http://localhost:8080/api/ordenes")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
