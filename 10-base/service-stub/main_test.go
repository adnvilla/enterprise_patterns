package main

import (
	"context"
	"errors"
	"testing"

	"github.com/adnvilla/enterprise_patterns/10-base/service-stub/pagos"
)

// Service Stub en acción: probamos el checkout — incluido el pago
// rechazado — sin tocar la pasarela real ni la red.
func TestCobrarPedidoConStub(t *testing.T) {
	ctx := context.Background()
	stub := &pagos.StubPaymentGateway{DeclineOver: 500000}

	// Camino feliz
	if err := cobrarPedido(ctx, stub, 12, 114897); err != nil {
		t.Fatalf("camino feliz: no esperaba error, obtuve %v", err)
	}
	if len(stub.Charges) != 1 || stub.Charges[0] != "stub-tx-12" {
		t.Fatalf("cobros registrados: obtuve %v, esperaba [stub-tx-12]", stub.Charges)
	}

	// Caso difícil bajo demanda: el stub rechaza los cobros enormes
	err := cobrarPedido(ctx, stub, 13, 999999)
	if !errors.Is(err, pagos.ErrPaymentDeclined) {
		t.Fatalf("esperaba ErrPaymentDeclined, obtuve %v", err)
	}
	if len(stub.Charges) != 1 {
		t.Fatalf("un cobro rechazado no debe registrarse: %v", stub.Charges)
	}
}
