package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/03-or-behavior/identity-map/pedidos"

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

	// Un finder por sesión/transacción de negocio
	finder := pedidos.NewOrderFinder(db)

	// Dos partes del código piden el mismo pedido...
	a, err := finder.Find(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	b, err := finder.Find(ctx, 12) // segunda llamada: NO toca la base de datos
	if err != nil {
		log.Fatal(err)
	}

	// ...y reciben la MISMA instancia
	fmt.Println("Pedido:", a.ID, "estado:", a.Status, "total:", a.TotalCents)
	fmt.Println("¿Misma instancia?", a == b) // true

	// Un cambio hecho vía `a` es visible vía `b`: no hay copias que se pisen
	a.Status = "paid"
	fmt.Println("Estado visto desde b:", b.Status) // "paid"
}
