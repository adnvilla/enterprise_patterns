package pagos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Cargo expresado en el vocabulario del dominio de la tienda
type Charge struct {
	OrderID     int64
	AmountCents int64
	Approved    bool
	Reference   string
}

// providerRequest y providerResponse son el formato PROPIO del proveedor:
// no salen de este paquete
type providerRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Ref      string `json:"external_reference"`
}

type providerResponse struct {
	Status string `json:"status"` // "ok" | "declined"
	TxID   string `json:"tx_id"`
}

// Gateway: encapsula la API HTTP del proveedor de pagos
type PaymentGateway struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewPaymentGateway(baseURL, apiKey string) *PaymentGateway {
	return &PaymentGateway{baseURL: baseURL, apiKey: apiKey, client: http.DefaultClient}
}

// La interfaz habla el idioma de la tienda, no el del proveedor
func (g *PaymentGateway) Charge(ctx context.Context, orderID, amountCents int64) (Charge, error) {
	req := providerRequest{
		Amount:   amountCents,
		Currency: "MXN",
		Ref:      fmt.Sprintf("order-%d", orderID),
	}
	body, err := json.Marshal(req)
	if err != nil {
		return Charge{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		g.baseURL+"/v1/charges", bytes.NewReader(body))
	if err != nil {
		return Charge{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+g.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(httpReq)
	if err != nil {
		return Charge{}, err
	}
	defer resp.Body.Close()

	var pr providerResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return Charge{}, err
	}
	return toCharge(orderID, amountCents, pr), nil
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
