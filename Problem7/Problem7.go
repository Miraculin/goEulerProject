package main

import (
	"fmt"
	"math"
)

func main() {

	curr := 3
	for n := 3; n <= 10001; n++ {
		curr += 2
		for !isPrime(curr) {
			curr += 2
		}
	}
	fmt.Println(curr)
}

func isPrime(i int) bool {
	max := int(math.Sqrt(float64(i))) + 1
	for j := 3; j < max; j += 2 {
		if i%j == 0 {
			return false
		}
	}
	return true
}
