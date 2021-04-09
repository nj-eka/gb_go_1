package prime

import (
	"fmt"
	"math"
)

func FindPrimesByBruteForce(n int) (primes []int) {
	primes = []int{2}
	fmt.Println(primes)
L:
	for i := 3; i <= n; i += 2 {
		fmt.Println(i)
		m := int(math.Floor(math.Sqrt(float64(i))))
		for j := 2; j <= m; j++ {
			if i%j == 0 {
				continue L
			}
		}
		primes = append(primes, i)
		fmt.Println(primes)
	}
	return
}

func FindPrimesBySieveOfEratosthenes(n int) (primes []int) {
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		fmt.Println(i)
		if b[i] {
			continue
		}
		primes = append(primes, i)
		fmt.Println(primes)
		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}
	return
}
