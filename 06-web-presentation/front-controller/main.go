package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/front-controller/front"
)

func main() {
	// La cadena completa: logging → auth → front controller → command.
	controller := front.New()
	app := front.WithLogging(front.WithAuth(controller))

	// Una sola entrada para todo el sitio.
	log.Println("Escuchando en :8080. Prueba con:")
	log.Println(`  curl -H "Authorization: Bearer demo" "http://localhost:8080/?cmd=orders.list"`)
	log.Println(`  curl -H "Authorization: Bearer demo" "http://localhost:8080/?cmd=orders.show&id=1"`)
	log.Println(`  curl "http://localhost:8080/?cmd=orders.list"   # sin auth → 401`)
	log.Fatal(http.ListenAndServe(":8080", app))
}
