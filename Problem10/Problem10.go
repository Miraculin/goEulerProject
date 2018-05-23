package main

import (
	"fmt"

	util "github.com/Miraculin/EulerProject"
)

func main() {
	sum := 2
	for n := 3; n <= 2000000; n += 2 {
		if util.IsPrime(n) {
			sum += n
		}
	}
	fmt.Println(sum)
}
