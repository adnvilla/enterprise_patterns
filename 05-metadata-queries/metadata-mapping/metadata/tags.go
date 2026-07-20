package metadata

import "reflect"

// Alternativa: los metadatos viven en struct tags, tal como hacen
// GORM (`gorm:"column:..."`) o sqlx (`db:"..."`).
type Order struct {
	ID         int64  `db:"id"`
	CustomerID int64  `db:"customer_id"`
	Status     string `db:"status"`
	TotalCents int64  `db:"total_cents"`
}

// MappingFromTags construye el TableMapping leyendo los tags `db`.
func MappingFromTags(table string, entity any) TableMapping {
	t := reflect.TypeOf(entity)
	m := TableMapping{Table: table}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if col := f.Tag.Get("db"); col != "" {
			m.Fields = append(m.Fields, FieldMapping{Field: f.Name, Column: col})
		}
	}
	return m
}
