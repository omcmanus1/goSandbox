package main

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/omcmanus1/goSandbox/internal/httpclient"
)

func main() {
	fmt.Printf("%# v\n", pretty.Formatter(httpclient.Resty()))
}
