package clientes

import "context"

// Supertype: lo que el resto del código necesita saber de un cliente
type Customer interface {
	Name() string
	DiscountPercent() int64 // % de descuento sobre el pedido
}

// Caso normal: un cliente real de la tabla customers
type RegisteredCustomer struct {
	ID    int64
	name  string
	email string
}

func (c *RegisteredCustomer) Name() string { return c.name }

func (c *RegisteredCustomer) DiscountPercent() int64 {
	return 5 // los clientes registrados tienen 5% en la tienda
}

// Special Case: el cliente desconocido, con valores SEGUROS
type UnknownCustomer struct{}

func (UnknownCustomer) Name() string           { return "invitado" }
func (UnknownCustomer) DiscountPercent() int64 { return 0 }

// Provider: el repositorio devuelve el caso especial en lugar de nil.
// (En la entrada la consulta va a la tabla customers con database/sql;
// aquí usamos un almacén en memoria para correr sin base de datos.)
type Repository struct {
	customers map[int64]*RegisteredCustomer
}

func NewRepository() *Repository {
	return &Repository{customers: map[int64]*RegisteredCustomer{
		1: {ID: 1, name: "Ana", email: "ana@example.com"},
	}}
}

func (r *Repository) ByID(ctx context.Context, id int64) (Customer, error) {
	c, ok := r.customers[id]
	if !ok {
		// Aquí está el patrón: ni nil, ni error.
		// "No encontrado" es un caso de negocio con comportamiento propio.
		return UnknownCustomer{}, nil
	}
	// Los errores REALES (red, SQL) sí se propagarían como error
	return c, nil
}
