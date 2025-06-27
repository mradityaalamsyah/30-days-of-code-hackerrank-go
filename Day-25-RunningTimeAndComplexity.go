package main

import (
	"fmt"
	"math"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	t := int(0)
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		n := int32(0)
		fmt.Scan(&n)

		if isPrime(n) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(n int32) bool {
	if n < 2 {
		return false
	}

	for i := int32(2); i <= int32(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
