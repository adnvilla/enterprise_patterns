package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/foreign-key-mapping/tienda"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://app:app@localhost:5432/tienda?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customers := &tienda.CustomerMapper{DB: db}

	// Uno-a-muchos vía Foreign Key Mapping: fila -> objeto.
	ana, err := customers.FindWithOrders(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s tiene %d pedidos\n", ana.Name, len(ana.Orders))

	// Y la dirección objeto -> fila: el Identity Field del padre
	// aterriza en la columna FK del hijo.
	nuevo := &tienda.Order{Status: "nuevo", TotalCents: 9900}
	if err := customers.AddOrder(ana, nuevo); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pedido %d creado; ahora %s tiene %d pedidos\n",
		nuevo.ID, ana.Name, len(ana.Orders))

	for _, o := range ana.Orders {
		fmt.Printf("  pedido %d: %s (%d centavos)\n", o.ID, o.Status, o.TotalCents)
	}
}
