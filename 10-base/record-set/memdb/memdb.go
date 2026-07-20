package memdb

import (
	"database/sql/driver"
	"errors"
	"io"
)

// Tabla simulada en memoria: columnas y filas en el orden de Columns.
type Table struct {
	Columns []string
	Rows    [][]driver.Value
}

// Driver mínimo de database/sql que sirve la tabla desde memoria.
// Es lo justo para que db.Query devuelva *sql.Rows REALES sin ninguna
// base de datos: el Record Set no distingue de dónde vienen las filas.
type Driver struct {
	Table Table
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return &conn{t: d.Table}, nil
}

type conn struct{ t Table }

func (c *conn) Prepare(query string) (driver.Stmt, error) { return &stmt{t: c.t}, nil }
func (c *conn) Close() error                              { return nil }
func (c *conn) Begin() (driver.Tx, error) {
	return nil, errors.New("transacciones no soportadas")
}

type stmt struct{ t Table }

func (s *stmt) Close() error  { return nil }
func (s *stmt) NumInput() int { return 0 }
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("solo consultas")
}
func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return &rows{t: s.t}, nil
}

type rows struct {
	t Table
	i int
}

func (r *rows) Columns() []string { return r.t.Columns }
func (r *rows) Close() error      { return nil }
func (r *rows) Next(dest []driver.Value) error {
	if r.i >= len(r.t.Rows) {
		return io.EOF
	}
	copy(dest, r.t.Rows[r.i])
	r.i++
	return nil
}
