package random

import (
	"log"

	"github.com/monaco-io/request"
)

func Request() string {

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
