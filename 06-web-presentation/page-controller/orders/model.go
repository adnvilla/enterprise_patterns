package orders

import "errors"

// Model: el dominio que ambos controladores consultan.
type Order struct {
	ID         int64
	Customer   string
	Status     string
	TotalCents int64
}

var ErrNotFound = errors.New("orden no encontrada")

type Repository struct{ data map[int64]Order }

func NewRepository() *Repository {
	return &Repository{data: map[int64]Order{
		1: {ID: 1, Customer: "Ana", Status: "pending", TotalCents: 149900},
		2: {ID: 2, Customer: "Luis", Status: "shipped", TotalCents: 89900},
	}}
}

func (r *Repository) All() []Order {
	list := make([]Order, 0, len(r.data))
	for _, o := range r.data {
		list = append(list, o)
	}
	return list
}

func (r *Repository) ByID(id int64) (Order, error) {
	o, ok := r.data[id]
	if !ok {
		return Order{}, ErrNotFound
	}
	return o, nil
}
