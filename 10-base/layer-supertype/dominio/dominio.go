package dominio

// Layer Supertype: el comportamiento común de toda la capa de dominio
type Entity struct {
	ID int64
}

// IsNew: aún no ha sido persistido
func (e Entity) IsNew() bool { return e.ID == 0 }

// En Go no heredamos: embebemos
type Order struct {
	Entity // Order "es un" objeto de dominio con ID
	CustomerID int64
	Status     string
	TotalCents int64
}

type Customer struct {
	Entity // mismo supertipo, cero duplicación
	Name  string
	Email string
}
