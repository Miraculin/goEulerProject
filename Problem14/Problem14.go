package main

import "fmt"

func main() {
	j := 0
	max_step := 0
	max_num := 1
	curr_step := 0
	for n := 1; n <= 1000000; n++ {
		j = n
		curr_step = 1
		for j != 1 {
			if j%2 == 0 {
				j = j / 2
				curr_step++
			} else {
				j = 3*j + 1
				curr_step++
			}
		}
		if curr_step > max_step {
			max_step = curr_step
			max_num = n
		}
	}
	fmt.Println(max_step, max_num)
}
