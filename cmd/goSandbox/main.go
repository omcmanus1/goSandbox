package main

import (
	"log"

	"github.com/omcmanus1/goSandbox/internal/webserver"
)

func main() {
	log.Fatal(webserver.CreateServer("3000"))
}
