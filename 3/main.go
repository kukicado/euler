/***
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
***/
package main

import "fmt"

func main() {
	n := 600851475143
	i := 2

	for i < n {
		for n%i == 0 {
			fmt.Println("The value of i is ", i)
			n = n / i
			fmt.Println("The value of n is ", n)
		}
		i++
	}
	fmt.Println(i)
}
