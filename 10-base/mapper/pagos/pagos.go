package pagos

import "encoding/json"

// Cargo expresado en el vocabulario del dominio de la tienda
type Charge struct {
	OrderID     int64
	AmountCents int64
	Approved    bool
	Reference   string
}

// providerResponse es el formato PROPIO del proveedor: no sale de este paquete
type providerResponse struct {
	Status string `json:"status"` // "ok" | "declined"
	TxID   string `json:"tx_id"`
}

// Mapper: traduce entre el modelo del proveedor y el del dominio.
// Ni Charge sabe que existe providerResponse, ni al revés.
func toCharge(orderID, amountCents int64, pr providerResponse) Charge {
	return Charge{
		OrderID:     orderID,
		AmountCents: amountCents,
		Approved:    pr.Status == "ok",
		Reference:   pr.TxID,
	}
}

// ProcessProviderReply es el "tercero" que invoca al mapper: recibe la
// respuesta cruda del proveedor, la decodifica a su modelo y deja que el
// mapper haga la traducción. Ninguno de los dos subsistemas — el dominio
// con su Charge, el proveedor con su JSON — sabe que la traducción ocurre.
func ProcessProviderReply(orderID, amountCents int64, raw []byte) (Charge, error) {
	var pr providerResponse
	if err := json.Unmarshal(raw, &pr); err != nil {
		return Charge{}, err
	}
	return toCharge(orderID, amountCents, pr), nil
}
