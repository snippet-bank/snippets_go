package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	jsonData := []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" },
						{ "id": "1003", "type": "Blueberry" },
						{ "id": "1004", "type": "Devil's Food" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5002", "type": "Glazed" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" },
				{ "id": "5003", "type": "Chocolate" },
				{ "id": "5004", "type": "Maple" }
			]
	}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})
	printJSON(data)
}

func printJSON(data interface{}) {
	printJSONIndented("", data, 0)
}

func printJSONIndented(key string, data interface{}, level int) {
	fmt.Printf(strings.Repeat("  ", level))

	switch data := data.(type) {
	case map[string]interface{}:
		fmt.Printf("[map]\n")
		for k, v := range data {
			printJSONIndented(k, v, level+1)
		}
	case string:
		fmt.Printf("%s: %s [string]\n", key, data)
	case float64:
		fmt.Printf("%s, %f [float]\n", key, data)
	case []interface{}:
		fmt.Println("[array]")
		for i, u := range data {
			printJSONIndented(strconv.Itoa(i), u, level+1)
		}
	default:
		fmt.Printf(key, data, "[unknown]\n")
	}
}
