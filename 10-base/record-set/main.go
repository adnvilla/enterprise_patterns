package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/adnvilla/enterprise_patterns/10-base/record-set/memdb"
	"github.com/adnvilla/enterprise_patterns/10-base/record-set/reportes"
)

func main() {
	// Sin Docker ni servidor: un driver mínimo registrado en database/sql
	// sirve filas simuladas en memoria — pero el *sql.Rows es real,
	// que es lo único que ToRecordSet necesita.
	sql.Register("memoria", &memdb.Driver{Table: memdb.Table{
		Columns: []string{"id", "customer_id", "status", "total_cents"},
		Rows: [][]driver.Value{
			{int64(12), int64(7), "paid", int64(114897)},
			{int64(13), int64(7), "pending", int64(999999)},
			{int64(14), int64(9), "paid", int64(34999)},
		},
	}})

	db, err := sql.Open("memoria", "tienda")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, customer_id, status, total_cents FROM orders`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// El reporte no necesita structs del dominio: filas y columnas genéricas
	rs, err := reportes.ToRecordSet(rows)
	if err != nil {
		log.Fatal(err)
	}

	// Como un grid: se recorre por nombre de columna, sin tipos propios
	for _, row := range rs {
		fmt.Printf("pedido %v del cliente %v: %v por %v centavos\n",
			row["id"], row["customer_id"], row["status"], row["total_cents"])
	}
	fmt.Println("Filas en el record set:", len(rs))
}
