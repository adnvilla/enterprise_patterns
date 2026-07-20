package orders

import (
	"html/template"
	"io"
)

// View: solo lee el modelo y lo pinta. Cero decisiones de negocio.
var listTemplate = template.Must(template.New("orders").Parse(`
<h1>Órdenes</h1>
<ul>
{{range .}}
  <li>#{{.ID}} — {{.Customer}} — ${{printf "%.2f" .TotalPesos}}
    {{if .CanBeCancelled}}<a href="/orders/cancel?id={{.ID}}">cancelar</a>{{end}}
  </li>
{{end}}
</ul>`))

func RenderList(w io.Writer, list []Order) error {
	return listTemplate.Execute(w, list)
}
