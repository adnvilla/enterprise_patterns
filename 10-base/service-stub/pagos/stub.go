package pagos

import (
	"context"
	"errors"
	"fmt"
)

var ErrPaymentDeclined = errors.New("pago rechazado por el emisor")

// Service Stub: implementación falsa del gateway, en memoria.
// Hoy lo llamaríamos "fake" o "test double"; la idea es la misma.
type StubPaymentGateway struct {
	// Configurable: montos que queremos que "fallen" en el test
	DeclineOver int64 // rechaza cobros mayores a este monto (0 = nunca)

	// Registro de lo cobrado, para poder afirmar en los tests
	Charges []string
}

func (s *StubPaymentGateway) Charge(ctx context.Context, orderID int64, amountCents int64) (string, error) {
	// Caso difícil bajo demanda: el gateway real no te deja
	// provocar un rechazo cuando tú quieras; el stub, sí.
	if s.DeclineOver > 0 && amountCents > s.DeclineOver {
		return "", ErrPaymentDeclined
	}

	txID := fmt.Sprintf("stub-tx-%d", orderID)
	s.Charges = append(s.Charges, txID)
	return txID, nil
}
