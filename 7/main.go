/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {

	i := 2
	primes := []int{}

	for len(primes) < 10001 {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Println(i, " is a Prime Number")
			primes = append(primes, i)
		}
		i++
	}

	fmt.Println(primes[10000])
}
