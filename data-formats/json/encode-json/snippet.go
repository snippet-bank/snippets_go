package snippets

import (
	"encoding/json"
	"log"
)

// User is either a Payer or a Payee.
type User struct {
	ID   string
	Name string
}

// Payment is a single transaction between a Payer and a Payee.
type Payment struct {
	ID     string
	Amount float64
	Payer  User
	Payee  User
}

func encodePayments(payments ...Payment) string {
	paymentsJSON, err := json.Marshal(payments)
	if err != nil {
		log.Fatal(err)
	}
	return string(paymentsJSON)
}
