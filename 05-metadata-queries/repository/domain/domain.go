package domain

// El dominio habla de pedidos, no de filas ni de SQL.
type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	TotalCents int64
}

// Repository: el dominio define QUÉ necesita; la infraestructura
// decidirá CÓMO. La interfaz vive AQUÍ, del lado del consumidor —
// puro idioma Go: «acepta interfaces, devuelve structs».
type OrderRepository interface {
	Add(o *Order) error
	PaidByCustomer(customerID int64) ([]Order, error)
}

// Lógica de dominio que usa el repositorio como si fuera una
// colección en memoria. Cero SQL a la vista.
func TotalSpent(repo OrderRepository, customerID int64) (int64, error) {
	orders, err := repo.PaidByCustomer(customerID)
	if err != nil {
		return 0, err
	}
	var total int64
	for _, o := range orders {
		total += o.TotalCents
	}
	return total, nil
}
