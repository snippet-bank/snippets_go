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

	want := `[{"ID":"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C","Amount":15.5,"Payer":{"ID":"363B8983-6751-4738-BDE1-C6E9C9B60729","Name":"Jack Black"},"Payee":{"ID":"566C6837-A02C-4932-8CAD-62DF35E50ED6","Name":"Joe Blow"}}]`
	got := encodePayments(payment)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
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

	got := encodePayments(payment1, payment2)
	want := `[{"ID":"79A3844F-A08F-4EE3-B4FA-F0C152D2EC0C","Amount":15.5,"Payer":{"ID":"363B8983-6751-4738-BDE1-C6E9C9B60729","Name":"Jack Black"},"Payee":{"ID":"566C6837-A02C-4932-8CAD-62DF35E50ED6","Name":"Joe Blow"}},{"ID":"DC3C23FE-7714-4A22-94D9-ABDA993A6FC1","Amount":15.5,"Payer":{"ID":"28E29D6F-A7FB-4C8B-9C49-702F1EEC731D","Name":"Jack Black"},"Payee":{"ID":"00A38362-59AE-4CFF-A07F-55E50E8EE355","Name":"Joe Blow"}}]`

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
