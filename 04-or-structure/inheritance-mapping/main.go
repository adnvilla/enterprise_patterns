package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/inheritance-mapping/catalogo"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// nuevosProductos arma un subtipo de cada sabor: los mismos dos objetos
// entran por las tres estrategias para poder compararlas.
func nuevosProductos() (catalogo.PhysicalProduct, catalogo.DigitalProduct) {
	teclado := catalogo.PhysicalProduct{WeightGrams: 900, Stock: 12}
	teclado.Name, teclado.PriceCents = "Teclado mecánico", 159900
	ebook := catalogo.DigitalProduct{DownloadURL: "https://cdn/tienda/go.epub",
		FileBytes: 4_500_000}
	ebook.Name, ebook.PriceCents = "eBook de Go", 29900
	return teclado, ebook
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://app:app@localhost:5432/tienda?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	teclado, ebook := nuevosProductos()

	// Estrategia A: Single Table Inheritance — una tabla, columna
	// discriminadora; al leer, cada fila regresa como su struct concreto.
	fmt.Println("— Single Table Inheritance —")
	single := &catalogo.SingleTableMapper{DB: db}
	idTeclado, err := single.Insert(teclado)
	if err != nil {
		log.Fatal(err)
	}
	idEbook, err := single.Insert(ebook)
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range []int64{idTeclado, idEbook} {
		p, err := single.FindByID(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p.Describe()) // polimorfismo vía interface
	}

	// Estrategia B: Class Table Inheritance — tabla por clase unidas por id;
	// insertar toca dos tablas y leer junta base + subtipo.
	fmt.Println("— Class Table Inheritance —")
	class := &catalogo.ClassTableMapper{DB: db}
	idTeclado, err = class.Insert(teclado)
	if err != nil {
		log.Fatal(err)
	}
	idEbook, err = class.Insert(ebook)
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range []int64{idTeclado, idEbook} {
		p, err := class.FindByID(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p.Describe())
	}

	// Estrategia C: Concrete Table Inheritance — tabla por clase concreta;
	// lo polimórfico se paga con un UNION.
	fmt.Println("— Concrete Table Inheritance —")
	concrete := &catalogo.ConcreteTableMapper{DB: db}
	if _, err := concrete.Insert(teclado); err != nil {
		log.Fatal(err)
	}
	if _, err := concrete.Insert(ebook); err != nil {
		log.Fatal(err)
	}
	all, err := concrete.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range all {
		fmt.Println(p.Describe())
	}
}
