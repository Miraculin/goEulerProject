package main

import "fmt"

func main() {
	ch := make(chan int)
	go fib(ch)
	sum := 0
	for v := range ch {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Println(sum)
}
func fib(c chan int) {
	for a, b := 0, 1; a < 4000000; a, b = a+b, a {
		c <- a
	}
	close(c)
}
