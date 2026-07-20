package query

// Query Object: la consulta es un objeto, no un string.
// Cada condición es un dato que luego se interpreta.
type condition struct {
	column string
	op     string
	value  any
}

type Query struct {
	table   string
	conds   []condition
	orderBy []string
	limit   int
}

func New(table string) *Query {
	return &Query{table: table}
}

// Where acumula una condición; no toca la base todavía.
func (q *Query) Where(column, op string, value any) *Query {
	q.conds = append(q.conds, condition{column: column, op: op, value: value})
	return q
}

func (q *Query) OrderBy(expr string) *Query {
	q.orderBy = append(q.orderBy, expr)
	return q
}

func (q *Query) Limit(n int) *Query {
	q.limit = n
	return q
}
