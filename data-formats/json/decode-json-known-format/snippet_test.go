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

	expected := "Alex Hamamatsu"
	payerName := pmts.Payments[0].Payer.Name
	if payerName != expected {
		t.Errorf("Got %s, expected %s ", payerName, expected)
	}

	expected = "Melba Miyamoto"
	payerName = pmts.Payments[1].Payer.Name
	if payerName != expected {
		t.Errorf("Got %s, expected %s ", payerName, expected)
	}

}
