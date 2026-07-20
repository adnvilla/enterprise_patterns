package metadata

// Metadata Mapping: el mapeo objeto-relacional vive en DATOS, no en código.
type FieldMapping struct {
	Field  string // campo del struct
	Column string // columna en la tabla
}

type TableMapping struct {
	Table  string
	Fields []FieldMapping
}

// El mapeo de Order es UNA declaración, no cuatro métodos con SQL repetido.
// (Idiomático en Go: un mapa explícito, sin magia.)
var OrderMapping = TableMapping{
	Table: "orders",
	Fields: []FieldMapping{
		{Field: "ID", Column: "id"},
		{Field: "CustomerID", Column: "customer_id"},
		{Field: "Status", Column: "status"},
		{Field: "TotalCents", Column: "total_cents"},
	},
}
