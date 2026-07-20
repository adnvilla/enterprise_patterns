package main

import (
	"fmt"

	"github.com/adnvilla/enterprise_patterns/10-base/layer-supertype/dominio"
)

func main() {
	// Un pedido recién creado: aún sin ID, aún sin persistir
	o := &dominio.Order{CustomerID: 7, Status: "pending", TotalCents: 114897}
	fmt.Println("¿Pedido nuevo?", o.IsNew()) // true: cortesía del supertipo

	// "Persistimos": la base de datos le asigna su ID
	o.ID = 42
	fmt.Println("¿Pedido nuevo tras guardar?", o.IsNew()) // false

	// Customer embebe el MISMO supertipo: cero duplicación en la capa
	c := &dominio.Customer{Name: "Ana", Email: "ana@example.com"}
	fmt.Println("¿Cliente nuevo?", c.IsNew()) // true

	c.ID = 7
	fmt.Printf("Cliente %s con id %d, pedido %d por %d centavos\n",
		c.Name, c.ID, o.ID, o.TotalCents)
}
