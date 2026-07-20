package main

import (
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/application-controller/checkout"
)

func main() {
	flow := checkout.NewFlowController()
	http.HandleFunc("/checkout", checkout.StepHandler(flow))
	log.Println("Checkout en http://localhost:8080/checkout")
	log.Println("Prueba con:")
	log.Println(`  curl "http://localhost:8080/checkout?state=cart&event=next"`)
	log.Println(`  curl "http://localhost:8080/checkout?state=payment&event=back"`)
	log.Println(`  curl "http://localhost:8080/checkout?state=cart&event=back"   # transición inválida → 400`)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
