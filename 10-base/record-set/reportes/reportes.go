package reportes

import "database/sql"

// Record Set: filas y columnas genéricas en memoria, sin structs.
// Así trabajaban el DataSet de ADO.NET y el RecordSet de ADO;
// en Go lo imitamos leyendo *sql.Rows hacia []map[string]any.
func ToRecordSet(rows *sql.Rows) ([]map[string]any, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var out []map[string]any
	for rows.Next() {
		values := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		row := make(map[string]any, len(cols))
		for i, c := range cols {
			row[c] = values[i] // columna -> valor, como una fila de grid
		}
		out = append(out, row)
	}
	return out, rows.Err()
}
