package catalogo

import (
	"database/sql"
	"fmt"
)

// SingleTableMapper organiza el mapeo de la jerarquía (Inheritance Mappers)
// para la estrategia A — Single Table Inheritance: toda la jerarquía en una
// tabla con columna discriminadora. La lógica común vive aquí; cada subtipo
// tiene su rama de carga.
type SingleTableMapper struct {
	DB *sql.DB
}

// FindByID lee la fila y usa la columna discriminadora
// para decidir qué struct concreto construir.
func (m *SingleTableMapper) FindByID(id int64) (Product, error) {
	var (
		typ           string
		base          baseProduct
		weight, stock sql.NullInt64
		downloadURL   sql.NullString
		fileBytes     sql.NullInt64
	)
	err := m.DB.QueryRow(
		`SELECT id, type, name, price_cents,
		        weight_grams, stock, download_url, file_bytes
		 FROM products WHERE id = $1`, id,
	).Scan(&base.ID, &typ, &base.Name, &base.PriceCents,
		&weight, &stock, &downloadURL, &fileBytes)
	if err != nil {
		return nil, err
	}

	switch typ { // el discriminador manda
	case "physical":
		return PhysicalProduct{baseProduct: base,
			WeightGrams: int(weight.Int64), Stock: int(stock.Int64)}, nil
	case "digital":
		return DigitalProduct{baseProduct: base,
			DownloadURL: downloadURL.String, FileBytes: fileBytes.Int64}, nil
	default:
		return nil, fmt.Errorf("tipo de producto desconocido: %q", typ)
	}
}

// Insert hace el camino inverso: del struct concreto a la fila plana.
func (m *SingleTableMapper) Insert(p Product) (int64, error) {
	var id int64
	switch v := p.(type) {
	case PhysicalProduct:
		err := m.DB.QueryRow(
			`INSERT INTO products (type, name, price_cents, weight_grams, stock)
			 VALUES ('physical', $1, $2, $3, $4) RETURNING id`,
			v.Name, v.PriceCents, v.WeightGrams, v.Stock).Scan(&id)
		return id, err
	case DigitalProduct:
		err := m.DB.QueryRow(
			`INSERT INTO products (type, name, price_cents, download_url, file_bytes)
			 VALUES ('digital', $1, $2, $3, $4) RETURNING id`,
			v.Name, v.PriceCents, v.DownloadURL, v.FileBytes).Scan(&id)
		return id, err
	default:
		return 0, fmt.Errorf("subtipo no mapeado: %T", p)
	}
}
