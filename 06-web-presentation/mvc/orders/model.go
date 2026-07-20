package orders

// Model: datos + lógica de dominio. No sabe nada de HTTP ni de HTML.
type Order struct {
	ID         int64
	Customer   string
	Status     string
	TotalCents int64
}

// Lógica de dominio: vive en el modelo, no en la plantilla ni en el handler.
func (o Order) CanBeCancelled() bool {
	return o.Status == "pending"
}

func (o Order) TotalPesos() float64 {
	return float64(o.TotalCents) / 100
}

// Repositorio mínimo para el ejemplo (en la serie ya vimos Repository).
type Repository struct{ data []Order }

func NewRepository() *Repository {
	return &Repository{data: []Order{
		{ID: 1, Customer: "Ana", Status: "pending", TotalCents: 149900},
		{ID: 2, Customer: "Luis", Status: "shipped", TotalCents: 89900},
	}}
}

func (r *Repository) All() []Order { return r.data }
