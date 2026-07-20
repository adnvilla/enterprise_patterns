package pagos

import "context"

// Separated Interface: el contrato es pequeño, y por eso
// escribir un stub es trivial. La definimos del lado del consumidor.
type PaymentGateway interface {
	// Charge cobra amount_cents al cliente y devuelve el id de la transacción
	Charge(ctx context.Context, orderID int64, amountCents int64) (string, error)
}
