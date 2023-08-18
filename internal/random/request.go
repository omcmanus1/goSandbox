package random

import (
	"log"

	"github.com/monaco-io/request"
)

func Request() string {

	type Item struct {
		ID          int                `json:"id"`
		Title       string             `json:"title"`
		Price       float32            `json:"price"`
		Category    string             `json:"category"`
		Description string             `json:"description"`
		Image       string             `json:"image"`
		Rating      map[string]float64 `json:"rating"`
	}

	var result []Item

	client := request.Client{
		URL:    "https://fakestoreapi.com/products",
		Method: "GET",
	}
	
	resp := client.Send().Scan(&result)
	if !resp.OK() {
		log.Println(resp.Error())
	}

	for _, item := range result {
		log.Println(item.Title)
	}

	responseString := resp.String()

	return responseString
}
