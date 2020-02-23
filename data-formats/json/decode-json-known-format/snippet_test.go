package snippets

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDecodePayment(t *testing.T) {
	jsonFile, err := os.Open("fixtures/payments.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	pmts := decodePayment(data)

	want := "Alex Hamamatsu"
	got := pmts.Payments[0].Payer.Name
	if got != want {
		t.Errorf("Got %s, want %s ", got, want)
	}

	want = "Melba Miyamoto"
	got = pmts.Payments[1].Payer.Name
	if got != want {
		t.Errorf("Got %s, want %s ", got, want)
	}

}
