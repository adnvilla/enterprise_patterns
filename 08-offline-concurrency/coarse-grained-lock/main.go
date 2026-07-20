package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/08-offline-concurrency/coarse-grained-lock/pedidos"
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

	// Dos usuarios leen el MISMO agregado (raíz + líneas, una sola versión)
	pedidoAna, err := repo.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	pedidoBeto, err := repo.FindByID(ctx, 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ana lee el pedido %d (version del agregado = %d, %d líneas)\n",
		pedidoAna.ID, pedidoAna.Version, len(pedidoAna.Lines))
	fmt.Printf("Beto lee el pedido %d (version del agregado = %d, %d líneas)\n",
		pedidoBeto.ID, pedidoBeto.Version, len(pedidoBeto.Lines))

	// Ana solo cambia UNA línea... pero el lock del agregado aplica igual
	pedidoAna.Lines[0].Quantity = 3
	if err := repo.Save(ctx, pedidoAna); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ana guarda una línea; nueva versión del agregado:", pedidoAna.Version)

	// Beto toca OTRA parte del agregado (la raíz), pero el candado es el mismo:
	// la versión de la raíz ya cambió y su guardado completo se aborta
	pedidoBeto.Status = "paid"
	if err := repo.Save(ctx, pedidoBeto); err != nil {
		if errors.Is(err, pedidos.ErrConflict) {
			fmt.Println("Beto: alguien más modificó este pedido. Recarga y vuelve a intentarlo.")
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Beto guarda el pedido; nueva versión del agregado:", pedidoBeto.Version)
}
