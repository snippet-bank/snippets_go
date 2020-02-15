package snippets

import (
	"testing"
)

func TestEncodeOnePayment(t *testing.T) {
	payment := Payment{
		"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C",
		15.50,
		User{
			"363B8983-6751-4738-BDE1-C6E9C9B60729",
			"Jack Black",
		},
		User{
			"566C6837-A02C-4932-8CAD-62DF35E50ED6",
			"Joe Blow",
		},
	}

	expected := `[{"id":"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C","amount":15.5,"payer":{"id":"363B8983-6751-4738-BDE1-C6E9C9B60729","name":"Jack Black"},"payee":{"id":"566C6837-A02C-4932-8CAD-62DF35E50ED6","name":"Joe Blow"}}]`
	json := encodePayments(payment)

	if json != expected {
		t.Errorf("Got %s, expected %s", json, expected)
	}
}

func TestEncodeMultiplePayments(t *testing.T) {
	payment1 := Payment{
		"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C",
		15.50,
		User{
			"363B8983-6751-4738-BDE1-C6E9C9B60729",
			"Jack Black",
		},
		User{
			"566C6837-A02C-4932-8CAD-62DF35E50ED6",
			"Joe Blow",
		},
	}

	payment2 := Payment{
		"DC3C23FE-7714-4A22-94D9-ABDA993A6FC1",
		15.50,
		User{
			"28E29D6F-A7FB-4C8B-9C49-702F1EEC731D",
			"Jack Black",
		},
		User{
			"00A38362-59AE-4CFF-A07F-55E50E8EE355",
			"Joe Blow",
		},
	}

	json := encodePayments(payment1, payment2)
	expected := `[{"id":"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C","amount":15.5,"payer":{"id":"363B8983-6751-4738-BDE1-C6E9C9B60729","name":"Jack Black"},"payee":{"id":"566C6837-A02C-4932-8CAD-62DF35E50ED6","name":"Joe Blow"}},{"id":"DC3C23FE-7714-4A22-94D9-ABDA993A6FC1","amount":15.5,"payer":{"id":"28E29D6F-A7FB-4C8B-9C49-702F1EEC731D","name":"Jack Black"},"payee":{"id":"00A38362-59AE-4CFF-A07F-55E50E8EE355","name":"Joe Blow"}}]`

	if json != expected {
		t.Errorf("Got %s, expected %s", json, expected)
	}
}
