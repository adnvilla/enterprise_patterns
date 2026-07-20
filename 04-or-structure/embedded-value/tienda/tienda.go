package tienda

import "database/sql"

// Address es un objeto de valor: existe en memoria,
// pero en la base son tres columnas de customers.
type Address struct {
	Street string
	City   string
	Zip    string
}

type Customer struct {
	ID      int64
	Name    string
	Email   string
	Address Address // Embedded Value
}

type CustomerMapper struct {
	DB *sql.DB
}

func (m *CustomerMapper) FindByID(id int64) (*Customer, error) {
	c := &Customer{}
	// El mapper aplana y des-aplana: columnas <-> campos del valor.
	err := m.DB.QueryRow(
		`SELECT id, name, email, address_street, address_city, address_zip
		 FROM customers WHERE id = $1`, id,
	).Scan(&c.ID, &c.Name, &c.Email,
		&c.Address.Street, &c.Address.City, &c.Address.Zip)
	if err != nil {
		return nil, err
	}
	return c, nil
}
