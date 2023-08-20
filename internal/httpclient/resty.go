package httpclient

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/omcmanus1/goSandbox/internal/types"
)

func Resty() types.Item {
	client := resty.New()
	var results []types.Item
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
