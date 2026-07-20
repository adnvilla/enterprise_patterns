package main

import (
	"fmt"

	"github.com/adnvilla/enterprise_patterns/10-base/value-object/dinero"
)

func main() {
	// El clásico que justifica todo esto: floats y dinero no se llevan
	a, b := 0.1, 0.2
	fmt.Println(a + b) // 0.30000000000000004

	// Precio del producto, como price_cents en la tabla products
	precio := dinero.New(34999, "MXN") // $349.99

	// Inmutabilidad: Multiply devuelve un Money NUEVO, el original no cambia
	total := precio.Multiply(3)
	fmt.Println("Precio unitario:", precio) // 349.99 MXN, intacto
	fmt.Println("Total de 3 unidades:", total)

	// Las reglas viven en el tipo: mezclar monedas es un error explícito
	if _, err := total.Add(dinero.New(500, "USD")); err != nil {
		fmt.Println("Rechazado:", err)
	}

	// Identidad por valor: dos Money con los mismos campos SON iguales
	fmt.Println(dinero.New(100, "MXN") == dinero.New(100, "MXN")) // true
}
