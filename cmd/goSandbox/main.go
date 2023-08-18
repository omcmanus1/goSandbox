package main

import (
	"fmt"

	"github.com/omcmanus1/goSandbox/internal/random"
)

func main() {
	fmt.Println(random.Request())
	// fmt.Println(random.ParseJSON(random.ApiRequest()))
}
