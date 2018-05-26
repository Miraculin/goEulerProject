package main

import "fmt"

func main() {
	//Using Euler's Formula
	a, b, c, ka, kb, kc := 0, 0, 0, 0, 0, 0
	limit := 1000
	m := 2
	for c < limit {
		for n := 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if a+b+c == 1000 {
				fmt.Println(a, b, c)
				fmt.Println(a * b * c)
				break
			}
			for k := 2; ka+kb+kc <= limit; k++ {
				ka, kb, kc = k*a, k*b, k*c
				if ka+kb+kc == 1000 {
					fmt.Println(ka, kb, kc)
					fmt.Println(a * b * c)
					break
				}
			}
		}
		m++
	}
}
