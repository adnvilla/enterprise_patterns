package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/02-data-source/row-data-gateway/pedidos"
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

	// Inserto una fila de apoyo (en la entrada del blog esto lo hace
	// el Table Data Gateway; aquí nos enfocamos en el Row Data Gateway).
	var id int64
	err = db.QueryRow(
		`INSERT INTO orders (customer_id, status, total_cents, version)
		 VALUES ($1, $2, $3, 1) RETURNING id`,
		42, "pending", 15990).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	// Con el Row Data Gateway hablo con UNA fila.
	order, err := pedidos.FindOrderRow(db, id)
	if err != nil {
		log.Fatal(err)
	}
	order.Status = "paid"
	if err := order.Update(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido pagado:", order.ID)

	// Y cuando ya no lo necesito, la misma fila sabe borrarse.
	if err := order.Delete(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido eliminado:", order.ID)
}
