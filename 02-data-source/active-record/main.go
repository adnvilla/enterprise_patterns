package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/02-data-source/active-record/pedidos"
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

	// Creo el pedido y él solito se inserta.
	order := pedidos.NewOrder(db, 42, 15990)
	if err := order.Insert(); err != nil {
		log.Fatal(err)
	}

	// Lo recupero, aplico la regla de negocio y él mismo se actualiza.
	order, err = pedidos.FindOrder(db, order.ID)
	if err != nil {
		log.Fatal(err)
	}
	if err := order.MarkAsPaid(); err != nil {
		log.Fatal(err)
	}
	if err := order.Update(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pedido %d ahora está %s\n", order.ID, order.Status)
}
