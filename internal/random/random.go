package random

import (
	"fmt"
	"math/rand"
	"github.com/fatih/color"
)

func Random() {
	textColour := color.New(color.FgHiGreen)
	output := fmt.Sprintf("Hello there for the %vth time", rand.Intn(20))
	textColour.Println(output)
	fmt.Printf("Hello there for the %vth time from random \n", rand.Intn(20))
}