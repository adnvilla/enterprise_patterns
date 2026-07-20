package query

import (
	"fmt"
	"strings"
)

// Lista blanca de operadores: nada de interpolar lo que llegue del usuario.
var allowedOps = map[string]bool{
	"=": true, "<>": true, "<": true, "<=": true,
	">": true, ">=": true, "LIKE": true,
}

// SQL traduce el Query Object a SQL con placeholders $1..$n.
// Aquí vive el «Interpreter» del patrón.
func (q *Query) SQL() (string, []any, error) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "SELECT * FROM %s", q.table)

	args := make([]any, 0, len(q.conds))
	for i, c := range q.conds {
		if !allowedOps[c.op] {
			return "", nil, fmt.Errorf("operador no permitido: %q", c.op)
		}
		if i == 0 {
			sb.WriteString(" WHERE ")
		} else {
			sb.WriteString(" AND ")
		}
		// El valor SIEMPRE viaja como parámetro, nunca concatenado.
		args = append(args, c.value)
		fmt.Fprintf(&sb, "%s %s $%d", c.column, c.op, len(args))
	}

	if len(q.orderBy) > 0 {
		sb.WriteString(" ORDER BY " + strings.Join(q.orderBy, ", "))
	}
	if q.limit > 0 {
		fmt.Fprintf(&sb, " LIMIT %d", q.limit)
	}
	return sb.String(), args, nil
}
