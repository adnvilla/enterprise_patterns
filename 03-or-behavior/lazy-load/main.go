package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/03-or-behavior/lazy-load/pedidos"

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

	ctx := context.Background()

	// Carga ligera: solo la fila de orders, sin líneas
	pedido, err := pedidos.FindOrder(ctx, db, 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedido:", pedido.ID, "estado:", pedido.Status)

	// Primer acceso: AQUÍ ocurre la consulta a order_lines
	lineas, err := pedido.Lines(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lineas {
		fmt.Printf("producto %d x %d\n", l.ProductID, l.Quantity)
	}

	// Segundo acceso: ya no toca la base de datos
	lineas, _ = pedido.Lines(ctx)
	fmt.Println("líneas:", len(lineas))
}
