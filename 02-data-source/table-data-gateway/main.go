package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/02-data-source/table-data-gateway/pedidos"
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

	// Con el Table Data Gateway hablo con la tabla completa.
	gateway := pedidos.NewOrdersGateway(db)
	id, err := gateway.Insert(42, "pending", 15990)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido insertado con id:", id)

	orders, err := gateway.FindByCustomer(42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("El cliente 42 tiene %d pedidos\n", len(orders))
	for _, o := range orders {
		fmt.Printf("  pedido %d: %s por %d centavos\n", o.ID, o.Status, o.TotalCents)
	}
}
