package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/10-base/plugin/registro"
)

func main() {
	// Plugin: la CONFIGURACIÓN decide, no la compilación
	modo := os.Getenv("PAYMENT_GATEWAY") // "real" o "sandbox"
	if modo == "" {
		modo = "sandbox"
	}
	fmt.Println("Implementación elegida por configuración:", modo)

	reg := registro.New()
	reg.Register("real", &registro.RealGateway{APIKey: os.Getenv("PAYMENT_API_KEY")})
	reg.Register("sandbox", &registro.SandboxGateway{})

	// El enlace ocurre UNA vez, al arrancar
	gw, err := reg.Gateway(modo)
	if err != nil {
		log.Fatal(err)
	}

	// De aquí en adelante, inyección explícita: el resto del código
	// recibe un PaymentGateway y nunca vuelve a tocar el registry
	if err := cobrar(context.Background(), gw); err != nil {
		log.Fatal(err)
	}
}

func cobrar(ctx context.Context, gw registro.PaymentGateway) error {
	return gw.Charge(ctx, 42, 129900)
}
