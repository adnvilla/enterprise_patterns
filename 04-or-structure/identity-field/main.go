package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/identity-field/pedidos"

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

	mapper := &pedidos.OrderMapper{DB: db}

	// Objeto en memoria, sin fila todavía: su id es el valor cero.
	order := &pedidos.Order{CustomerID: 7, Status: "nuevo", TotalCents: 25900}
	fmt.Println("¿Es nuevo?", order.IsNew()) // true

	if err := mapper.Insert(order); err != nil {
		log.Fatal(err)
	}
	// A partir de aquí, objeto y fila comparten identidad.
	fmt.Println("Pedido guardado con id:", order.ID)
	fmt.Println("¿Es nuevo?", order.IsNew()) // false

	same, err := mapper.FindByID(order.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Recuperado:", same.ID, same.Status)
}
