package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/association-table-mapping/tienda"

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

	products := &tienda.ProductMapper{DB: db}

	// Muchos-a-muchos vía Association Table Mapping.
	if err := products.ReplaceTags(42, []int64{1, 3}); err != nil { // oferta, gadgets
		log.Fatal(err)
	}
	tags, err := products.FindTags(42)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tags {
		fmt.Println("tag:", t.Name)
	}
}
