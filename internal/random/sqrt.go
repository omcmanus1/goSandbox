package random

import "fmt"

func Sqrt(x float64) (string, float64) {
	z := float64(1)
	prevGuess := float64(0)
	count := 1
	for count < 20 {
		prevGuess = z
		z -= (z*z - x) / (2 * z)
		if prevGuess == z {
			fmt.Println(count)
			return "early answer: ", prevGuess
		} else {
			count++
		}
	}
	fmt.Println(count)
	return "final answer: ", z
}

func SquarerRoot(x float64) (string, float64) {
	z := float64(1)
	prevGuess := float64(0)
	count := 1
	for z != prevGuess {
		prevGuess = z
		z -= (z*z - x) / (2 * z)
		count++
	}
	fmt.Println(count)
	return "answer: ", z
}
