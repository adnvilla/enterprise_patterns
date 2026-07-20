package dinero

import "testing"

func TestAllocate(t *testing.T) {
	// El ejemplo clásico del libro: 10 centavos entre 3 → 4, 3 y 3
	parts := New(10, "MXN").Allocate(3)
	want := []int64{4, 3, 3}
	for i, p := range parts {
		if p.Amount != want[i] {
			t.Errorf("parte %d: obtuve %d centavos, esperaba %d", i, p.Amount, want[i])
		}
		if p.Currency != "MXN" {
			t.Errorf("parte %d: la moneda debe conservarse, obtuve %q", i, p.Currency)
		}
	}

	// 1148.97 en 2 pagos: 574.49 y 574.48 — el centavo huérfano no se pierde
	pagos := New(114897, "MXN").Allocate(2)
	if pagos[0].Amount != 57449 || pagos[1].Amount != 57448 {
		t.Errorf("obtuve %d y %d, esperaba 57449 y 57448", pagos[0].Amount, pagos[1].Amount)
	}

	// Nunca se pierde (ni aparece) un centavo: la suma reconstruye el total
	var suma int64
	for _, p := range pagos {
		suma += p.Amount
	}
	if suma != 114897 {
		t.Errorf("la suma de las partes es %d, esperaba 114897", suma)
	}
}
