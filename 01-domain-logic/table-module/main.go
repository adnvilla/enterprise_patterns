package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/01-domain-logic/table-module/orders"
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

	// Un solo módulo para toda la tabla: se crea una vez
	// y sirve para cualquier cliente y cualquier pedido.
	module := orders.NewOrdersModule(db)

	total, err := module.TotalByCustomer(42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total histórico del cliente 42:", total, "centavos")

	overdue, err := module.MarkOverdue(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedidos marcados como vencidos:", overdue)
}
