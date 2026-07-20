package checkout

import (
	"fmt"
	"net/http"
)

// Input Controller delgado: NO sabe de reglas de navegación.
func StepHandler(flow *FlowController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// En la vida real, el estado actual saldría de la sesión.
		current := State(r.FormValue("state"))
		ev := Event(r.FormValue("event"))

		next, err := flow.Advance(current, ev)
		if err != nil {
			http.Error(w, "no puedes saltarte pasos del checkout", http.StatusBadRequest)
			return
		}

		// Aquí guardarías `next` en la sesión y renderizarías la vista
		// (con el Template View de hace dos entradas, por ejemplo).
		fmt.Fprintf(w, "siguiente pantalla: %s (%s)", next, flow.ViewFor(next))
	}
}
