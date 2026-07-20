package front

import (
	"fmt"
	"net/http"
)

// Command: la unidad de trabajo a la que el Front Controller despacha.
type Command interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

// Un command por acción; equivalen a los Page Controllers, pero
// viven DETRÁS de la puerta única.
type ListOrders struct{}

func (ListOrders) Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "listado de órdenes")
}

type ShowOrder struct{}

func (ShowOrder) Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "detalle de la orden %s\n", r.URL.Query().Get("id"))
}
