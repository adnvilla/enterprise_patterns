package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/dependent-mapping/tienda"

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

	orders := &tienda.OrderMapper{DB: db}

	// Dependent Mapping: las líneas viajan con su pedido, siempre.
	order := &tienda.Order{ID: 12, Status: "pagado", TotalCents: 51800,
		Lines: []tienda.OrderLine{{ProductID: 42, Quantity: 2}}}
	if err := orders.Update(order); err != nil {
		log.Fatal(err)
	}

	// Al releer, los dependientes regresan junto con su dueño.
	same, err := orders.FindByID(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pedido %d (%s) con %d línea(s):\n", same.ID, same.Status, len(same.Lines))
	for _, l := range same.Lines {
		fmt.Printf("  producto %d x %d\n", l.ProductID, l.Quantity)
	}
}
