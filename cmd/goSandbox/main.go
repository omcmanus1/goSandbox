package main

import (
	"fmt"

	"github.com/omcmanus1/goSandbox/internal/random"
)

func printMain() {
	fmt.Println("Hi from main")
}

func main() {
	printMain()
	random.Random()
}
