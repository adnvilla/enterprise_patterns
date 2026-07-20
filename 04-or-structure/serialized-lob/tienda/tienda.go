package tienda

import (
	"database/sql"
	"encoding/json"
)

// Specs es un grafo pequeño y de forma variable:
// perfecto para serializarse completo.
type Specs struct {
	Color      string            `json:"color,omitempty"`
	WeightG    int               `json:"weight_g,omitempty"`
	Dimensions map[string]int    `json:"dimensions,omitempty"`
	Extra      map[string]string `json:"extra,omitempty"`
}

type ProductMapper struct {
	DB *sql.DB
}

func (m *ProductMapper) SaveSpecs(productID int64, s Specs) error {
	blob, err := json.Marshal(s) // objeto -> LOB
	if err != nil {
		return err
	}
	_, err = m.DB.Exec(
		`UPDATE products SET specs = $1 WHERE id = $2`, blob, productID)
	return err
}

func (m *ProductMapper) FindSpecs(productID int64) (Specs, error) {
	var blob []byte
	var s Specs
	err := m.DB.QueryRow(
		`SELECT COALESCE(specs, '{}') FROM products WHERE id = $1`,
		productID).Scan(&blob)
	if err != nil {
		return s, err
	}
	return s, json.Unmarshal(blob, &s) // LOB -> objeto
}
