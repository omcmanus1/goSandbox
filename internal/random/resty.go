package random

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Resty() Item {
	client := resty.New()
	var results []Item
	resp, err := client.
		R().
		// SetResult(results).
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Oops, error")
	}
	if err := json.Unmarshal(resp.Body(), &results); err != nil {
		fmt.Println("Error:", err)
	}
	return results[4]
}
