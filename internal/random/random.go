package random

import (
	"fmt"
	"math/rand"
)

func Random() {
	output := fmt.Sprintf("Hello there for the %vth time", rand.Intn(20))
	fmt.Println(output)
	fmt.Printf("Hello there for the %vth time from random \n", rand.Intn(20))
}