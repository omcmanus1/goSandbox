package httpclient

import (
	"log"

	"github.com/monaco-io/request"
	"github.com/omcmanus1/goSandbox/internal/types"
)

func Request() string {

	var result []types.Item

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
