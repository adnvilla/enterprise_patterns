package registro

import (
	"context"
	"fmt"
)

// Gateway "real": le pega al proveedor de verdad
type RealGateway struct {
	APIKey string
}

func (g *RealGateway) Charge(ctx context.Context, orderID, amountCents int64) error {
	// aquí iría la llamada HTTP real al proveedor
	fmt.Printf("[REAL] cobrando %d centavos del pedido %d\n", amountCents, orderID)
	return nil
}

// Gateway "sandbox": aprueba todo, para desarrollo y staging
type SandboxGateway struct{}

func (g *SandboxGateway) Charge(ctx context.Context, orderID, amountCents int64) error {
	fmt.Printf("[SANDBOX] cobro simulado de %d centavos del pedido %d\n", amountCents, orderID)
	return nil
}
