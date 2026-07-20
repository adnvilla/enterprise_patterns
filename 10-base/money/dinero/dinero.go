package dinero

import (
	"errors"
	"fmt"
)

// Value Object: identidad por valor, inmutable, comparable con ==
// Amount va en centavos (int64), igual que price_cents y total_cents
// en las tablas products y orders. Nunca floats.
type Money struct {
	Amount   int64  // centavos
	Currency string // "MXN", "USD"...
}

var ErrCurrencyMismatch = errors.New("no se pueden operar montos de monedas distintas")

func New(amount int64, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}

// Add devuelve un Money NUEVO: los value objects no se modifican.
// Mezclar monedas es un error explícito, no un bug silencioso.
func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, ErrCurrencyMismatch
	}
	return Money{Amount: m.Amount + other.Amount, Currency: m.Currency}, nil
}

// Multiply escala el monto (precio x cantidad) sin salir de enteros
func (m Money) Multiply(qty int64) Money {
	return Money{Amount: m.Amount * qty, Currency: m.Currency}
}

// Allocate reparte el monto en n partes SIN perder centavos:
// el residuo se distribuye de a un centavo entre las primeras partes.
// Es el ejemplo clásico del libro: 10 entre 3 da 4, 3 y 3 — no 3.33 x 3.
func (m Money) Allocate(n int) []Money {
	base := m.Amount / int64(n)
	remainder := m.Amount % int64(n)

	parts := make([]Money, n)
	for i := range parts {
		parts[i] = Money{Amount: base, Currency: m.Currency}
		if int64(i) < remainder {
			parts[i].Amount++ // un centavo extra a las primeras partes
		}
	}
	return parts
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%02d %s", m.Amount/100, m.Amount%100, m.Currency)
}
