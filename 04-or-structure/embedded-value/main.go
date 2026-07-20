package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/embedded-value/tienda"

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

	// Embedded Value: la dirección llega ya armada como objeto.
	ana, err := customers.FindByID(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enviar a:", ana.Address.Street, ana.Address.City)
	fmt.Println("CP:", ana.Address.Zip)
}
