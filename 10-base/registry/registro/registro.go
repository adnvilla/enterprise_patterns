package registro

import (
	"context"
	"fmt"
	"sync"
)

// El contrato común de todos los gateways de pago
type PaymentGateway interface {
	Charge(ctx context.Context, orderID, amountCents int64) error
}

// Registry: objeto bien conocido donde se encuentran los servicios.
// Explícito y pasado por constructor — NO una variable global.
type Registry struct {
	mu       sync.RWMutex
	gateways map[string]PaymentGateway
}

func New() *Registry {
	return &Registry{gateways: make(map[string]PaymentGateway)}
}

func (r *Registry) Register(name string, g PaymentGateway) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.gateways[name] = g
}

func (r *Registry) Gateway(name string) (PaymentGateway, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	g, ok := r.gateways[name]
	if !ok {
		return nil, fmt.Errorf("gateway de pagos no registrado: %q", name)
	}
	return g, nil
}
