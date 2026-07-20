package front

import "net/http"

// Front Controller: TODA petición entra por aquí.
type FrontController struct {
	commands map[string]Command // dispatcher: nombre de comando → Command
}

func New() *FrontController {
	return &FrontController{commands: map[string]Command{
		"orders.list": ListOrders{},
		"orders.show": ShowOrder{},
	}}
}

func (f *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Despacho: aquí decidimos por comando; igual de válido sería por ruta.
	cmd, ok := f.commands[r.URL.Query().Get("cmd")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	cmd.Execute(w, r)
}
