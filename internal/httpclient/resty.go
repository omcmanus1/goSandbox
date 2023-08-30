package httpclient

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/kr/pretty"
	"github.com/omcmanus1/goSandbox/internal/types"
)

func Resty() string {
	client := resty.New()
	var results []types.Item
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Oops, error")
	}
	if err := json.Unmarshal(resp.Body(), &results); err != nil {
		fmt.Println("Error:", err)
	}
	return fmt.Sprintf("%# v \n", pretty.Formatter(results))
}
