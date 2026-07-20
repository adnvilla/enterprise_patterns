package vistas

import (
	"embed"
	"html/template"
	"net/http"
	"time"
)

//go:embed templates/*.tmpl
var files embed.FS

// Template Processor: se parsea UNA vez al arrancar, se ejecuta por request.
var ordersTmpl = template.Must(
	template.New("orders.tmpl").Funcs(Funcs).ParseFS(files, "templates/orders.tmpl"),
)

// Handler: prepara el view model y delega el pintado a la plantilla.
func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	page := OrdersPage{
		Customer: "Ana",
		Orders: []OrderView{ // normalmente vendrían del repositorio
			{ID: 1, Status: "PAID", TotalCents: 129900, CreatedAt: time.Now()},
			{ID: 2, Status: "PENDING", TotalCents: 45000, CreatedAt: time.Now()},
		},
	}
	if err := ordersTmpl.Execute(w, page); err != nil {
		http.Error(w, "error al renderizar", http.StatusInternalServerError)
	}
}
