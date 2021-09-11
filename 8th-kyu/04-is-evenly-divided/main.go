/* DIVIDE WATERMELONS INTO EVEN PIECES FOR 2 PEOPLE */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(DivideSolution1(8))
	fmt.Println(DivideSolution1(67))
	fmt.Println(DivideSolution2(2))
	fmt.Println(DivideSolution2(5))
}

func DivideSolution1(weight int) bool {
	if weight <= 2 {
		return false
	}
	return weight%2 == 0
}

func DivideSolution2(weight int) bool {
	if weight <= 2 {
		return false
	}

	// check for even number
	return (weight & 1) == 0
}
