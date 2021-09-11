/* FIND WHICH QUARTER OF THE YEAR THE MONTH IS IN */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(QuarterOfSolution1(8))
	fmt.Println(QuarterOfSolution1(10))
	fmt.Println(QuarterOfSolution2(2))
	fmt.Println(QuarterOfSolution2(5))
}

func QuarterOfSolution1(month int) int {
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	}

	return 0
}

func QuarterOfSolution2(month int) int {
	return (month + 2) / 3
}
