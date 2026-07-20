package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/special-case/clientes"
)

func main() {
	repo := clientes.NewRepository()
	ctx := context.Background()

	// El id 999 no existe en customers... y no pasa nada.
	customer, err := repo.ByID(ctx, 999)
	if err != nil {
		log.Fatal(err) // solo errores reales llegan aquí
	}

	// Cero ifs: el checkout trata igual al invitado y al registrado.
	// Antes: if customer == nil { name = "invitado" } else { ... }
	fmt.Printf("Hola, %s\n", customer.Name()) // "Hola, invitado"

	totalCents := int64(114897) // total_cents del pedido
	descuento := totalCents * customer.DiscountPercent() / 100
	fmt.Println("Descuento aplicado:", descuento) // 0, sin panic ni if

	// El mismo código, con un cliente registrado
	customer, err = repo.ByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hola, %s\n", customer.Name())
	descuento = totalCents * customer.DiscountPercent() / 100
	fmt.Println("Descuento aplicado:", descuento) // 5744
}
