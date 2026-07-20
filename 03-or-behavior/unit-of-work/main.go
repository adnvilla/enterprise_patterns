package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/03-or-behavior/unit-of-work/pedidos"

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

	uow := pedidos.NewUnitOfWork(db)

	// Una transacción de negocio que toca varios pedidos
	nuevo := &pedidos.Order{CustomerID: 7, Status: "pending", TotalCents: 45900}
	uow.RegisterNew(nuevo)

	pagado := &pedidos.Order{ID: 12, CustomerID: 3, Status: "paid", TotalCents: 129900}
	uow.RegisterDirty(pagado)

	cancelado := &pedidos.Order{ID: 15}
	uow.RegisterRemoved(cancelado)

	// Todo se escribe junto: o entra todo, o no entra nada
	if err := uow.Commit(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido nuevo con id:", nuevo.ID)

	// Consultamos el estado final para ver el resultado del commit
	rows, err := db.QueryContext(context.Background(),
		`SELECT id, customer_id, status, total_cents, version
		 FROM orders ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Pedidos en la base de datos:")
	for rows.Next() {
		var o pedidos.Order
		var version int64
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.Status, &o.TotalCents, &version); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  pedido %d: cliente %d, estado %s, total %d, versión %d\n",
			o.ID, o.CustomerID, o.Status, o.TotalCents, version)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
