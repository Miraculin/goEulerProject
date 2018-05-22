package main

import "fmt"

func main() {
	res := 0
	for i := 2520; ; i += 2520 {
		for j := 11; j < 21; j++ {
			if i%j != 0 {
				break
			}
			if j == 20 {
				res = i
				break
			}
		}
		fmt.Println(i)
		if res == i {
			break
		}
	}
	fmt.Println(res)
}
