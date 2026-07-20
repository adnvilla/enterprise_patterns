package tienda

import "database/sql"

type Tag struct {
	ID   int64
	Name string
}

type ProductMapper struct {
	DB *sql.DB
}

// FindTags lee fila -> objeto atravesando la tabla de asociación con JOIN.
// Nota que product_tags nunca se convierte en struct.
func (m *ProductMapper) FindTags(productID int64) ([]Tag, error) {
	rows, err := m.DB.Query(
		`SELECT t.id, t.name
		 FROM tags t
		 JOIN product_tags pt ON pt.tag_id = t.id
		 WHERE pt.product_id = $1
		 ORDER BY t.name`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var t Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, rows.Err()
}

// ReplaceTags escribe objeto -> fila: la forma más simple y robusta
// es borrar las asociaciones viejas e insertar las actuales.
func (m *ProductMapper) ReplaceTags(productID int64, tagIDs []int64) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(
		`DELETE FROM product_tags WHERE product_id = $1`, productID); err != nil {
		return err
	}
	for _, tagID := range tagIDs {
		if _, err := tx.Exec(
			`INSERT INTO product_tags (product_id, tag_id) VALUES ($1, $2)`,
			productID, tagID); err != nil {
			return err
		}
	}
	return tx.Commit()
}
