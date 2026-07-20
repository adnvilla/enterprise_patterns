package orders

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

// Page Controller del listado: SOLO sabe atender GET /orders.
type OrdersHandler struct {
	Repo *Repository
}

var listTpl = template.Must(template.New("list").Parse(
	`<h1>Órdenes</h1><ul>{{range .}}<li><a href="/orders/detail?id={{.ID}}">#{{.ID}} — {{.Customer}}</a></li>{{end}}</ul>`))

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Este controlador decide su modelo (todas las órdenes)...
	list := h.Repo.All()
	// ...y su plantilla (el listado).
	listTpl.Execute(w, list)
}

// Page Controller del detalle: SOLO sabe atender GET /orders/detail.
type OrderDetailHandler struct {
	Repo *Repository
}

var detailTpl = template.Must(template.New("detail").Parse(
	`<h1>Orden #{{.ID}}</h1><p>Cliente: {{.Customer}}</p><p>Estado: {{.Status}}</p>`))

func (h *OrderDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// La lógica de entrada de ESTA página: parsear y validar su parámetro.
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}
	o, err := h.Repo.ByID(id)
	if errors.Is(err, ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	detailTpl.Execute(w, o)
}
