package main

import "fmt"

func main() {
	nums := []int{}
	for i := 3; i < 1000; i += 3 {
		if i%5 != 0 {
			nums = append(nums, i)
		}
	}
	for i := 5; i < 1000; i += 5 {
		nums = append(nums, i)
	}
	sum := 0
	for _, v := range nums {
		sum += v
		fmt.Println(v)
	}
	fmt.Println(sum)
}
