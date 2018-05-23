package eulerProjectUtil

import "math"

func IsPrime(i int) bool {
	if i%2 == 0 {
		return false
	}
	max := int(math.Sqrt(float64(i))) + 1
	for j := 3; j < max; j += 2 {
		if i%j == 0 {
			return false
		}
	}
	return true
}
