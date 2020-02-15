package snippets

import (
	"encoding/json"
	"log"
)

// User is either a Payer or a Payee.
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Payment is a single transaction between a Payer and a Payee.
type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Payer  User    `json:"payer"`
	Payee  User    `json:"payee"`
}

func encodePayments(payments ...Payment) string {
	paymentsJSON, err := json.Marshal(payments)
	if err != nil {
		log.Fatal(err)
	}
	return string(paymentsJSON)
}
