package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/01-domain-logic/transaction-script/orders"
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

	// El cliente solo invoca el procedimiento:
	// una transacción de negocio, una función.
	orderID, err := orders.PlaceOrder(db, 42, []orders.Line{
		{ProductID: 1, Quantity: 2},
		{ProductID: 3, Quantity: 1},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido creado:", orderID)

	// Consultamos lo que el script dejó en la base de datos.
	var total int64
	if err := db.QueryRow(
		`SELECT total_cents FROM orders WHERE id = $1`, orderID,
	).Scan(&total); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total guardado:", total, "centavos")
}
