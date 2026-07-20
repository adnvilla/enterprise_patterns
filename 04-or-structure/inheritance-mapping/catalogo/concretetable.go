package catalogo

import (
	"database/sql"
	"fmt"
)

// ConcreteTableMapper implementa la estrategia C — Concrete Table Inheritance:
// una tabla por clase concreta, cada una con TODAS sus columnas.
// Cada subtipo es autocontenido; lo polimórfico cuesta un UNION.
type ConcreteTableMapper struct {
	DB *sql.DB
}

// Insert va directo a la tabla del subtipo: sin JOINs ni tabla base.
func (m *ConcreteTableMapper) Insert(p Product) (int64, error) {
	var id int64
	switch v := p.(type) {
	case PhysicalProduct:
		err := m.DB.QueryRow(
			`INSERT INTO physical_products_c (name, price_cents, weight_grams, stock)
			 VALUES ($1, $2, $3, $4) RETURNING id`,
			v.Name, v.PriceCents, v.WeightGrams, v.Stock).Scan(&id)
		return id, err
	case DigitalProduct:
		err := m.DB.QueryRow(
			`INSERT INTO digital_products_c (name, price_cents, download_url, file_bytes)
			 VALUES ($1, $2, $3, $4) RETURNING id`,
			v.Name, v.PriceCents, v.DownloadURL, v.FileBytes).Scan(&id)
		return id, err
	default:
		return 0, fmt.Errorf("subtipo no mapeado: %T", p)
	}
}

// FindAll muestra el precio de esta estrategia: la consulta polimórfica
// requiere un UNION entre las tablas concretas.
func (m *ConcreteTableMapper) FindAll() ([]Product, error) {
	rows, err := m.DB.Query(
		`SELECT 'physical' AS type, id, name, price_cents,
		        weight_grams, stock, '' AS download_url, 0 AS file_bytes
		 FROM physical_products_c
		 UNION ALL
		 SELECT 'digital', id, name, price_cents,
		        0, 0, download_url, file_bytes
		 FROM digital_products_c
		 ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var (
			typ           string
			base          baseProduct
			weight, stock int
			downloadURL   string
			fileBytes     int64
		)
		if err := rows.Scan(&typ, &base.ID, &base.Name, &base.PriceCents,
			&weight, &stock, &downloadURL, &fileBytes); err != nil {
			return nil, err
		}
		switch typ {
		case "physical":
			products = append(products, PhysicalProduct{baseProduct: base,
				WeightGrams: weight, Stock: stock})
		case "digital":
			products = append(products, DigitalProduct{baseProduct: base,
				DownloadURL: downloadURL, FileBytes: fileBytes})
		default:
			return nil, fmt.Errorf("tipo de producto desconocido: %q", typ)
		}
	}
	return products, rows.Err()
}
