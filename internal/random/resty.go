package random

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Resty() Item {
	client := resty.New()
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Oops, error")
	} else {
		fmt.Println("Response JSON:")
	}
	var results []Item
	if err := json.Unmarshal(resp.Body(), &results); err != nil {
		fmt.Println("Error:", err)
	}
	return results[2]
}
