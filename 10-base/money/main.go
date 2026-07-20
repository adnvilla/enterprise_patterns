package main

import (
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/money/dinero"
)

func main() {
	// El clásico que justifica todo esto: floats y dinero no se llevan
	a, b := 0.1, 0.2
	fmt.Println(a + b) // 0.30000000000000004

	// Precio del producto, como price_cents en la tabla products
	precio := dinero.New(34999, "MXN") // $349.99

	// total_cents del pedido: 3 unidades más el envío
	total := precio.Multiply(3) // 1049.97 MXN
	envio := dinero.New(9900, "MXN")

	total, err := total.Add(envio)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total del pedido:", total) // 1148.97 MXN

	// Mezclar monedas es un error explícito
	if _, err := total.Add(dinero.New(500, "USD")); err != nil {
		fmt.Println("Rechazado:", err)
	}

	// Repartir el total en 2 pagos sin perder centavos
	for i, pago := range total.Allocate(2) {
		fmt.Printf("Pago %d: %s\n", i+1, pago) // 574.49 y 574.48
	}

	// Identidad por valor: dos Money con los mismos campos SON iguales
	fmt.Println(dinero.New(100, "MXN") == dinero.New(100, "MXN")) // true
}
