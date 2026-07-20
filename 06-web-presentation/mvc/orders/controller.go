package orders

import "net/http"

// Controller: recibe la petición, consulta al modelo y elige la vista.
type ListController struct {
	Repo *Repository
}

func (c *ListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list := c.Repo.All() // habla con el modelo
	if err := RenderList(w, list); err != nil { // delega en la vista
		http.Error(w, "error al renderizar", http.StatusInternalServerError)
	}
}
