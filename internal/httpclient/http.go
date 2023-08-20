package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/omcmanus1/goSandbox/internal/types"
	"github.com/kr/pretty"
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
