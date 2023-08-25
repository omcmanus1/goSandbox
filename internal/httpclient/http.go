package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kr/pretty"
	"github.com/omcmanus1/goSandbox/internal/types"
)

func ApiRequest() []byte {
	res, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		log.Fatal("didn't work: ", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func ParseJSON(data []byte) string {
	var message []types.Item
	if err := json.Unmarshal([]byte(data), &message); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Printf("%# v \n", pretty.Formatter(message[4]))
	return message[5].Title
}

func ParsedRequest() interface{} {
	var data []types.Item

	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Error fetching products: ", err)
		return pretty.Formatter(data)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding product data: ", err)
		return pretty.Formatter(data)
	}

	return fmt.Sprintf("%# v \n", pretty.Formatter(data))
}
