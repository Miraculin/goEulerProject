package main

import (
	"fmt"

	util "github.com/Miraculin/EulerProject"
)

func main() {

	curr := 3
	for n := 3; n <= 10001; n++ {
		curr += 2
		for !util.IsPrime(curr) {
			curr += 2
		}
	}
	fmt.Println(curr)
}
