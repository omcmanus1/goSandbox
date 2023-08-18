package random

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RestyItem struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Price       float32            `json:"price"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Rating      map[string]float64 `json:"rating"`
}

func Resty() RestyItem {
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
	var results []RestyItem
	if err := json.Unmarshal(resp.Body(), &results); err != nil {
		fmt.Println("Error:", err)
	}
	return results[2]
}
