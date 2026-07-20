package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/08-offline-concurrency/optimistic-offline-lock/pedidos"
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

	repo := pedidos.NewOrderRepository(db)
	ctx := context.Background()

	// Dos transacciones de negocio leen el MISMO pedido con la MISMA versión
	pedidoAna, err := repo.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	pedidoBeto, err := repo.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ana lee el pedido %d (version = %d)\n", pedidoAna.ID, pedidoAna.Version)
	fmt.Printf("Beto lee el pedido %d (version = %d)\n", pedidoBeto.ID, pedidoBeto.Version)

	// ... ambos editan en pantalla durante diez minutos ...
	pedidoAna.Status = "paid"
	pedidoBeto.Status = "cancelled"

	// Ana confirma primero: su versión coincide y el UPDATE condicional gana
	if err := repo.Update(ctx, pedidoAna); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ana guarda el pedido; nueva versión:", pedidoAna.Version)

	// Beto confirma después: el UPDATE condicional detecta que alguien se adelantó
	if err := repo.Update(ctx, pedidoBeto); err != nil {
		if errors.Is(err, pedidos.ErrConflict) {
			// La capa de arriba lo traduce a lenguaje humano
			fmt.Println("Beto: alguien más modificó este pedido. Recarga y vuelve a intentarlo.")
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Beto guarda el pedido; nueva versión:", pedidoBeto.Version)
}
