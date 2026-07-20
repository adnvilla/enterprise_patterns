package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/08-offline-concurrency/pessimistic-offline-lock/pedidos"
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

	locks := pedidos.NewLockManager(db)
	ctx := context.Background()
	recurso := pedidos.OrderResource(12)

	// Identidad de cada transacción de negocio (sesión de usuario)
	ana := "sesion-ana-8f3a"
	beto := "sesion-beto-2c71"

	// ANTES de editar: Ana reserva el pedido
	if err := locks.Acquire(ctx, recurso, ana); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ana adquiere el lock de %q y abre el formulario.\n", recurso)

	// Beto intenta reservar el MISMO pedido: la clave primaria lo rechaza
	if err := locks.Acquire(ctx, recurso, beto); err != nil {
		if errors.Is(err, pedidos.ErrLockNotAcquired) {
			fmt.Println("Beto: este pedido lo está editando alguien más. Intenta más tarde.")
		} else {
			log.Fatal(err)
		}
	}

	// ... Ana edita con calma: el pedido es suyo ...
	if _, err := db.ExecContext(ctx,
		`UPDATE orders SET status = $1, version = version + 1 WHERE id = $2`,
		"paid", 12); err != nil {
		log.Fatal(err)
	}
	// Al confirmar, soltar el lock
	if err := locks.Release(ctx, recurso, ana); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ana guarda el pedido y libera el lock.")

	// Con el lock libre, ahora sí Beto puede reservar el pedido
	if err := locks.Acquire(ctx, recurso, beto); err != nil {
		log.Fatal(err)
	}
	// Pase lo que pase, soltar el lock al terminar
	defer locks.Release(ctx, recurso, beto)
	fmt.Printf("Beto adquiere el lock de %q ahora que está libre.\n", recurso)
}
