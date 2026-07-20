package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adnvilla/enterprise_patterns/04-or-structure/serialized-lob/tienda"

	_ "github.com/jackc/pgx/v5/stdlib"
)

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

	products := &tienda.ProductMapper{DB: db}

	// Serialized LOB: el grafo completo entra y sale de una columna JSONB.
	err = products.SaveSpecs(42, tienda.Specs{
		Color: "negro", WeightG: 350,
		Dimensions: map[string]int{"alto_mm": 120, "ancho_mm": 60},
	})
	if err != nil {
		log.Fatal(err)
	}

	specs, err := products.FindSpecs(42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Color:", specs.Color)
	fmt.Println("Peso (g):", specs.WeightG)
	fmt.Println("Dimensiones:", specs.Dimensions)
}
