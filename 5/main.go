/* 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import "fmt"

func main() {

	i := 20
	result := false

	for result == false {
		divisible := isDivisible(i)
		if divisible {
			fmt.Println(i)
			result = true
		} else {
			i = i + 20
		}

	}

}

func isDivisible(num int) bool {
	if num%20 == 0 && num%19 == 0 && num%18 == 0 && num%17 == 0 && num%16 == 0 && num%15 == 0 && num%14 == 0 && num%13 == 0 && num%12 == 0 && num%11 == 0 {
		return true
	}
	return false
}
