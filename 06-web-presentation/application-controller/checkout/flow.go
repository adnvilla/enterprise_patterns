package checkout

import "errors"

// Estados del flujo de compra.
type State string

const (
	Cart         State = "cart"
	Address      State = "address"
	Payment      State = "payment"
	Confirmation State = "confirmation"
)

// Eventos que puede disparar el usuario.
type Event string

const (
	Next Event = "next"
	Back Event = "back"
)

var ErrInvalidTransition = errors.New("transición no permitida")

// Application Controller: el mapa del flujo vive en UN solo lugar.
type FlowController struct {
	transitions map[State]map[Event]State
	views       map[State]string // estado -> vista a mostrar
}

func NewFlowController() *FlowController {
	return &FlowController{
		transitions: map[State]map[Event]State{
			Cart:    {Next: Address},
			Address: {Next: Payment, Back: Cart},
			Payment: {Next: Confirmation, Back: Address},
			// Confirmation es terminal: de ahí no hay salida.
		},
		views: map[State]string{
			Cart:         "cart.tmpl",
			Address:      "address.tmpl",
			Payment:      "payment.tmpl",
			Confirmation: "confirmation.tmpl",
		},
	}
}

// Advance aplica las reglas de navegación.
func (f *FlowController) Advance(current State, ev Event) (State, error) {
	next, ok := f.transitions[current][ev]
	if !ok {
		return current, ErrInvalidTransition
	}
	return next, nil
}

// ViewFor devuelve la vista que corresponde a un estado.
func (f *FlowController) ViewFor(s State) string {
	return f.views[s]
}
