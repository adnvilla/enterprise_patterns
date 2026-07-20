package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/adnvilla/enterprise_patterns/10-base/service-stub/pagos"
)

// El checkout depende de la INTERFAZ, no de la pasarela real
func cobrarPedido(ctx context.Context, gw pagos.PaymentGateway, orderID, totalCents int64) error {
	txID, err := gw.Charge(ctx, orderID, totalCents)
	if err != nil {
		return fmt.Errorf("cobrando pedido %d: %w", orderID, err)
	}
	fmt.Println("Pedido cobrado, transacción:", txID)
	return nil
}

func main() {
	ctx := context.Background()

	// En producción inyectaríamos el gateway real (Stripe, PayPal...).
	// En desarrollo y tests, el stub — el cliente ni se entera.
	stub := &pagos.StubPaymentGateway{DeclineOver: 500000}

	// Camino feliz: total_cents del pedido 12
	_ = cobrarPedido(ctx, stub, 12, 114897) // "stub-tx-12"

	// Caso difícil, provocado a voluntad: un cobro enorme
	err := cobrarPedido(ctx, stub, 13, 999999)
	fmt.Println("¿Rechazado?", errors.Is(err, pagos.ErrPaymentDeclined)) // true

	fmt.Println("Cobros registrados:", stub.Charges) // [stub-tx-12]
}
