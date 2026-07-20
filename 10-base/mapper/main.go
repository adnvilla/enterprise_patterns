package main

import (
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/mapper/pagos"
)

func main() {
	// Respuestas crudas tal como las mandaría el proveedor de pagos:
	// su formato ("status", "tx_id") jamás sale del paquete pagos
	aprobada := []byte(`{"status": "ok", "tx_id": "tx-8891"}`)
	rechazada := []byte(`{"status": "declined", "tx_id": ""}`)

	// El mapper trabaja invisible: aquí solo vemos entrar JSON del
	// proveedor y salir un Charge del dominio
	charge, err := pagos.ProcessProviderReply(42, 129900, aprobada)
	if err != nil {
		log.Fatal(err)
	}
	if charge.Approved {
		fmt.Println("Pedido cobrado, referencia:", charge.Reference)
	}

	charge, err = pagos.ProcessProviderReply(43, 55000, rechazada)
	if err != nil {
		log.Fatal(err)
	}
	if !charge.Approved {
		fmt.Println("Pago rechazado para el pedido", charge.OrderID)
	}
}
