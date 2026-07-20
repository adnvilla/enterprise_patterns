package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/query-object/query"
)

// PaidOrders habla el idioma del dominio, no el de la base:
// «los pedidos pagados de este cliente, los más caros primero».
func PaidOrders(customerID int64) *query.Query {
	return query.New("orders").
		Where("customer_id", "=", customerID).
		Where("status", "=", "paid").
		OrderBy("total_cents DESC").
		Limit(10)
}

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

	// Unos pedidos de ejemplo para que la consulta tenga qué encontrar.
	for _, o := range []struct {
		status string
		total  int64
	}{
		{"paid", 149900},
		{"paid", 25900},
		{"pending", 99900},
	} {
		if _, err := db.Exec(
			`INSERT INTO orders (customer_id, status, total_cents, version)
			 VALUES ($1, $2, $3, 1)`,
			42, o.status, o.total,
		); err != nil {
			log.Fatal(err)
		}
	}

	sqlText, args, err := PaidOrders(42).SQL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sqlText)
	// SELECT * FROM orders WHERE customer_id = $1 AND status = $2
	//   ORDER BY total_cents DESC LIMIT 10

	rows, err := db.Query(sqlText, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Escanear filas como de costumbre (orders: id, customer_id, status, total_cents, version).
	for rows.Next() {
		var (
			id, customerID, totalCents, version int64
			status                              string
		)
		if err := rows.Scan(&id, &customerID, &status, &totalCents, &version); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  pedido %d: cliente %d, %s por %d centavos\n",
			id, customerID, status, totalCents)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
