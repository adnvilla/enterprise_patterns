package vistas

import (
	"fmt"
	"html/template"
	"time"
)

// View model: datos ya listos para pintar. La plantilla no calcula nada.
type OrdersPage struct {
	Customer string
	Orders   []OrderView
}

type OrderView struct {
	ID         int64
	Status     string
	TotalCents int64
	CreatedAt  time.Time
}

// Helpers de presentación: viven en Go, donde se pueden probar.
var Funcs = template.FuncMap{
	// Convierte centavos a moneda legible.
	"money": func(cents int64) string {
		return fmt.Sprintf("$%.2f", float64(cents)/100)
	},
	// Formatea fechas al estilo de la casa.
	"fecha": func(t time.Time) string {
		return t.Format("02/01/2006")
	},
}
