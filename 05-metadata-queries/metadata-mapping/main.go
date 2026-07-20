package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/metadata-mapping/metadata"
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

	order := metadata.Order{ID: 1, CustomerID: 42, Status: "paid", TotalCents: 149900}

	// Los metadatos salen de los tags; el SQL sale de los metadatos.
	m := metadata.MappingFromTags("orders", order)

	fmt.Println(m.InsertSQL())
	// INSERT INTO orders (id, customer_id, status, total_cents)
	//   VALUES ($1, $2, $3, $4)

	if _, err := db.Exec(m.InsertSQL(), m.Values(order)...); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.SelectSQL())
	// SELECT id, customer_id, status, total_cents FROM orders WHERE id = $1

	// Leemos de vuelta el pedido usando el SELECT generado desde los metadatos.
	var got metadata.Order
	if err := db.QueryRow(m.SelectSQL(), order.ID).
		Scan(&got.ID, &got.CustomerID, &got.Status, &got.TotalCents); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pedido leído: id=%d customer=%d status=%s total=%d centavos\n",
		got.ID, got.CustomerID, got.Status, got.TotalCents)
}
