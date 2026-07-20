package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/domain"
	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/postgres"
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

	// El concreto se elige en el borde de la aplicación (el main);
	// de aquí para adentro, todo el mundo ve solo la interfaz.
	var repo domain.OrderRepository = &postgres.OrderRepository{DB: db}

	if err := repo.Add(&domain.Order{
		CustomerID: 42, Status: "paid", TotalCents: 149900,
	}); err != nil {
		log.Fatal(err)
	}

	total, err := domain.TotalSpent(repo, 42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("El cliente 42 ha gastado %d centavos\n", total)
}
