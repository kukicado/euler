/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var largestPalindrome int
	var multiple1 int
	var multiple2 int

	for first := 999; first > 99; first-- {
		for second := 999; second > 99; second-- {
			value := strconv.Itoa(first * second)
			// "335523"
			isPalindrome := false
			for i := 0; i < len(value)/2; i++ {
				if value[i] == value[len(value)-i-1] {
					isPalindrome = true
				} else {
					isPalindrome = false
					break
				}
			}

			if isPalindrome {
				palindrome, _ := strconv.Atoi(value)
				if palindrome > largestPalindrome {
					multiple1 = first
					multiple2 = second
					largestPalindrome = palindrome
				}
			}

		}
	}
	fmt.Println(multiple1)
	fmt.Println(multiple2)
	fmt.Println(largestPalindrome)
}
