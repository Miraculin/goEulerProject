package main

import (
	"fmt"
	"math"
)

func main() {
	n := 600851475143
	res := n
	max := int(math.Sqrt(float64(n))) + 1
	for i := 2; i <= max; i++ {
		if i == n {
			res = i
		}
		if n%i == 0 {
			n = n / i
			max = int(math.Sqrt(float64(n))) + 1
		}
	}
	if n > max {
		fmt.Println(n)
	} else {
		fmt.Println(res)
	}
}
