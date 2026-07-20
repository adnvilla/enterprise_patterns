package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/template-view/vistas"
)

func main() {
	http.HandleFunc("/ordenes", vistas.OrdersHandler)
	log.Println("Escuchando en http://localhost:8080/ordenes")
	log.Println("Prueba con:")
	log.Println("  curl http://localhost:8080/ordenes")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
