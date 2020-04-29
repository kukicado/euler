/*The sum of the squares of the first ten natural numbers is,

12+22+...+102=385
The square of the sum of the first ten natural numbers is,

(1+2+...+10)2=552=3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025−385=2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package main

import "fmt"

func main() {
	ints, squares := getSums(500)
	difference := ints - squares
	fmt.Println(ints, squares, difference)
}

func getSums(n int) (int, int) {
	sumOfSquares := ((n * (n + 1)) * (2*n + 1)) / 6
	sumOfInts := (n * (n + 1)) / 2
	sumOfInts = sumOfInts * sumOfInts

	return sumOfInts, sumOfSquares
}
