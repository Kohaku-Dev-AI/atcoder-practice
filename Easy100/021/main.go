package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)

	for {
		if isPrime(x) {
			fmt.Println(x)
			break
		}
		x++
	}
}
