package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestIntegerFinderSolution1([]int{34, 15, 88, 2}))
	fmt.Println(SmallestIntegerFinderSolution2([]int{34, -345, -1, 100}))
}

func SmallestIntegerFinderSolution1(numbers []int) int {
	smallest := numbers[0]

	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
}

func SmallestIntegerFinderSolution2(numbers []int) int {
	sort.Ints(numbers)

	return numbers[0]
}
