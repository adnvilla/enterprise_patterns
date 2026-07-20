package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/registry/registro"
)

func main() {
	// Registry: se crea en main y se pasa explícitamente — no es global
	reg := registro.New()
	reg.Register("real", &registro.RealGateway{APIKey: "sk_live_123"})
	reg.Register("sandbox", &registro.SandboxGateway{})

	// Punto de encuentro bien conocido: los servicios se buscan por nombre
	gw, err := reg.Gateway("sandbox")
	if err != nil {
		log.Fatal(err)
	}

	// El enlace ocurre UNA vez; de aquí en adelante, inyección explícita:
	// cobrar recibe el gateway ya resuelto, no el registry
	if err := cobrar(context.Background(), gw); err != nil {
		log.Fatal(err)
	}

	// Buscar un servicio no registrado falla en runtime, no en compilación
	if _, err := reg.Gateway("paypal"); err != nil {
		fmt.Println("Error esperado:", err)
	}
}

func cobrar(ctx context.Context, gw registro.PaymentGateway) error {
	return gw.Charge(ctx, 42, 129900)
}
