package domain_test

import (
	"testing"

	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/domain"
	"github.com/adnvilla/enterprise_patterns/05-metadata-queries/repository/memory"
)

// El test de la lógica de dominio corre en microsegundos,
// sin Docker, sin Postgres, sin fixtures de SQL.
func TestTotalSpent(t *testing.T) {
	repo := &memory.OrderRepository{}
	repo.Add(&domain.Order{CustomerID: 42, Status: "paid", TotalCents: 149900})
	repo.Add(&domain.Order{CustomerID: 42, Status: "pending", TotalCents: 99900})

	total, err := domain.TotalSpent(repo, 42)
	if err != nil {
		t.Fatal(err)
	}
	if total != 149900 { // el pedido pendiente no cuenta
		t.Fatalf("esperaba 149900, obtuve %d", total)
	}
}
