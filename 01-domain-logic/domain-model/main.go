package main

import (
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/01-domain-logic/domain-model/domain"
)

func main() {
	// La lógica rica en acción: el cliente conversa con el negocio,
	// no con tablas ni con SQL.
	order := domain.NewOrder(42)

	if err := order.AddLine(1, 2, 4500); err != nil { // 2 playeras
		log.Fatal(err)
	}
	if err := order.AddLine(3, 1, 12000); err != nil { // 1 sudadera
		log.Fatal(err)
	}

	order.ApplyDiscount() // el pedido decide si el descuento aplica

	if err := order.Confirm(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pedido %s por %d centavos\n", order.Status, order.Total())

	// Las invariantes se defienden solas:
	if err := order.AddLine(5, 1, 900); err != nil {
		fmt.Println("Rechazado:", err) // pedido confirmado, no acepta líneas
	}
}
