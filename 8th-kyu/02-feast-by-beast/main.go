package main

import "fmt"

func main() {
	fmt.Println(FeastSolution1("brown bear", "blue car"))
	fmt.Println(FeastSolution2("brown bear", "blue car"))
}

func FeastSolution1(beast string, dish string) bool {
	beastArr := [2]string{string(beast[0]), string(beast[len(beast)-1])}

	dishArr := [2]string{string(dish[0]), string(dish[len(dish)-1])}

	return beastArr == dishArr
}

func FeastSolution2(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}
