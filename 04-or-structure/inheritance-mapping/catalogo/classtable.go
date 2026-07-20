package catalogo

import (
	"database/sql"
	"fmt"
)

// ClassTableMapper implementa la estrategia B — Class Table Inheritance:
// una tabla por clase (base incluida), unidas por el mismo id.
// Escribir toca dos tablas; leer requiere JOIN con la tabla del subtipo.
type ClassTableMapper struct {
	DB *sql.DB
}

// Insert reparte el struct concreto entre la tabla base y la del subtipo,
// dentro de una transacción para que compartan id de forma atómica.
func (m *ClassTableMapper) Insert(p Product) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int64
	switch v := p.(type) {
	case PhysicalProduct:
		if err := tx.QueryRow(
			`INSERT INTO products_base (name, price_cents)
			 VALUES ($1, $2) RETURNING id`,
			v.Name, v.PriceCents).Scan(&id); err != nil {
			return 0, err
		}
		if _, err := tx.Exec(
			`INSERT INTO physical_products (id, weight_grams, stock)
			 VALUES ($1, $2, $3)`,
			id, v.WeightGrams, v.Stock); err != nil {
			return 0, err
		}
	case DigitalProduct:
		if err := tx.QueryRow(
			`INSERT INTO products_base (name, price_cents)
			 VALUES ($1, $2) RETURNING id`,
			v.Name, v.PriceCents).Scan(&id); err != nil {
			return 0, err
		}
		if _, err := tx.Exec(
			`INSERT INTO digital_products (id, download_url, file_bytes)
			 VALUES ($1, $2, $3)`,
			id, v.DownloadURL, v.FileBytes); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("subtipo no mapeado: %T", p)
	}
	return id, tx.Commit()
}

// FindByID lee la fila base y luego busca en qué tabla de subtipo vive el id:
// la tabla donde aparece la fila hace de discriminador.
func (m *ClassTableMapper) FindByID(id int64) (Product, error) {
	var base baseProduct
	err := m.DB.QueryRow(
		`SELECT id, name, price_cents FROM products_base WHERE id = $1`, id,
	).Scan(&base.ID, &base.Name, &base.PriceCents)
	if err != nil {
		return nil, err
	}

	var weightGrams, stock int
	err = m.DB.QueryRow(
		`SELECT weight_grams, stock FROM physical_products WHERE id = $1`, id,
	).Scan(&weightGrams, &stock)
	if err == nil {
		return PhysicalProduct{baseProduct: base,
			WeightGrams: weightGrams, Stock: stock}, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	var downloadURL string
	var fileBytes int64
	err = m.DB.QueryRow(
		`SELECT download_url, file_bytes FROM digital_products WHERE id = $1`, id,
	).Scan(&downloadURL, &fileBytes)
	if err != nil {
		return nil, err
	}
	return DigitalProduct{baseProduct: base,
		DownloadURL: downloadURL, FileBytes: fileBytes}, nil
}
