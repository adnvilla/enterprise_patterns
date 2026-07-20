package metadata

import (
	"fmt"
	"reflect"
	"strings"
)

// SelectSQL genera el SELECT a partir de los metadatos.
func (m TableMapping) SelectSQL() string {
	cols := make([]string, len(m.Fields))
	for i, f := range m.Fields {
		cols[i] = f.Column
	}
	return fmt.Sprintf("SELECT %s FROM %s WHERE id = $1",
		strings.Join(cols, ", "), m.Table)
}

// InsertSQL genera el INSERT con placeholders $1..$n estilo Postgres.
func (m TableMapping) InsertSQL() string {
	cols := make([]string, len(m.Fields))
	marks := make([]string, len(m.Fields))
	for i, f := range m.Fields {
		cols[i] = f.Column
		marks[i] = fmt.Sprintf("$%d", i+1)
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		m.Table, strings.Join(cols, ", "), strings.Join(marks, ", "))
}

// Values extrae del struct, vía reflection, los valores en el orden
// de los metadatos. Este es el «programa reflexivo» de Fowler.
func (m TableMapping) Values(entity any) []any {
	v := reflect.ValueOf(entity)
	out := make([]any, len(m.Fields))
	for i, f := range m.Fields {
		out[i] = v.FieldByName(f.Field).Interface()
	}
	return out
}
