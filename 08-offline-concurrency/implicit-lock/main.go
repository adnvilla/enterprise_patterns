package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/08-offline-concurrency/implicit-lock/pedidos"
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

	// Cada usuario trabaja con SU repositorio (su sesión / unit of work).
	// Fíjate que el main no menciona versiones ni locks en ningún momento.
	repoAna := pedidos.NewOrderRepository(db)
	repoBeto := pedidos.NewOrderRepository(db)

	pedidoAna, err := repoAna.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	pedidoBeto, err := repoBeto.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ana y Beto leen el pedido 12; el repositorio adquirió el lock por ellos.")

	// Ana edita y guarda: el chequeo de versión ocurre solo, dentro de Save
	pedidoAna.Status = "paid"
	if err := repoAna.Save(ctx, pedidoAna); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ana guarda el pedido sin haber escrito una sola línea de locking.")

	// Beto guarda después: el lock que él tampoco pidió lo protege igual
	pedidoBeto.TotalCents = 39900
	if err := repoBeto.Save(ctx, pedidoBeto); err != nil {
		if errors.Is(err, pedidos.ErrConflict) {
			fmt.Println("Beto: alguien más modificó este pedido. Recarga y vuelve a intentarlo.")
			fmt.Println("(El conflicto lo detectó el repositorio automáticamente: imposible olvidarlo.)")
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Beto guarda el pedido.")
}
