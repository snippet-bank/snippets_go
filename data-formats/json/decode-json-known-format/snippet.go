package snippets

import "encoding/json"

// Payments is a collection of payments.
type Payments struct {
	Payments []Payment `json:"payments"`
}

// Payment is one successful transfer of funds.
type Payment struct {
	ID       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"` // TODO: can we make this an enum?
	Comment  string  `json:"comment"`
	Payer    User    `json:"payer"`
	Payee    User    `json:"payee"`
}

// User is a minimal user representation,
// just enough to display or process a list of payments.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func decodePayment(payload []byte) *Payments {
	var payments Payments

	json.Unmarshal(payload, &payments)

	return &payments
}
