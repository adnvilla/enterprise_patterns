package webview

// Paso 1: representación lógica de la pantalla.
// Dice QUÉ se muestra, todavía no CÓMO se ve.
type Screen struct {
	Title   string
	Section string
	Orders  []OrderRow
}

type OrderRow struct {
	ID    int64
	Total string
}

// Cada página del sitio construye su pantalla lógica.
func OrdersScreen(rows []OrderRow) Screen {
	return Screen{
		Title:   "Tus órdenes",
		Section: "cuenta",
		Orders:  rows,
	}
}
