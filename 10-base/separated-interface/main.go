package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/separated-interface/dominio"
	"github.com/adnvilla/enterprise_patterns/10-base/separated-interface/memoria"
)

// El servicio depende de la INTERFAZ del dominio, no de la infraestructura
func cobrarPedido(ctx context.Context, repo dominio.OrderRepository, id int64) error {
	o, err := repo.ByID(ctx, id)
	if err != nil {
		return err
	}
	o.Status = "paid"
	return repo.Save(ctx, o)
}

func main() {
	ctx := context.Background()

	// La implementación concreta se conecta hasta el main.
	// En producción sería postgres.NewOrderRepository(db) — mismo contrato.
	var repo dominio.OrderRepository = memoria.NewOrderRepository()

	// Sembramos un pedido pendiente
	pedido := &dominio.Order{CustomerID: 7, Status: "pending", TotalCents: 114897}
	if err := repo.Save(ctx, pedido); err != nil {
		log.Fatal(err)
	}

	if err := cobrarPedido(ctx, repo, pedido.ID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pedido %d cobrado\n", pedido.ID)

	cobrado, err := repo.ByID(ctx, pedido.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Estado final: %s (%d centavos)\n", cobrado.Status, cobrado.TotalCents)
}
