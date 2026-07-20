package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/02-data-source/data-mapper/dominio"
	"github.com/adnvilla/enterprise_patterns/02-data-source/data-mapper/persistencia"
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

	mapper := persistencia.NewOrderMapper(db)

	// El dominio se crea y opera sin saber de la base.
	order := dominio.NewOrder(42, 15990)
	if err := order.ApplyDiscount(10); err != nil {
		log.Fatal(err)
	}
	if err := mapper.Insert(order); err != nil {
		log.Fatal(err)
	}

	// Recupero, aplico negocio y persisto: cada quien en su carril.
	order, err = mapper.Find(order.ID)
	if err != nil {
		log.Fatal(err)
	}
	if err := order.MarkAsPaid(); err != nil {
		log.Fatal(err)
	}
	if err := mapper.Update(order); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pedido %d: %s por %d centavos\n", order.ID, order.Status, order.TotalCents)
}
