package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/adnvilla/enterprise_patterns/06-web-presentation/two-step-view/webview"
)

//go:embed templates/*.tmpl
var files embed.FS

var tmpl = template.Must(template.ParseFS(files, "templates/*.tmpl"))

func main() {
	// Two Step View: pantalla lógica -> layout -> HTML.
	http.HandleFunc("/ordenes", func(w http.ResponseWriter, r *http.Request) {
		screen := webview.OrdersScreen([]webview.OrderRow{ // paso 1
			{ID: 1, Total: "$1,299.00"},
			{ID: 2, Total: "$450.00"},
		})
		if err := tmpl.ExecuteTemplate(w, "layout", screen); err != nil { // paso 2
			http.Error(w, "error al renderizar", http.StatusInternalServerError)
		}
	})

	log.Println("Escuchando en :8080. Prueba con:")
	log.Println("  curl http://localhost:8080/ordenes")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
