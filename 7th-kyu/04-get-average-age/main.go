package main

import "fmt"

type Developer struct {
	firstName, lastName, country, continent, language string
	age                                               int
}

var list1 = []Developer{
	{
		firstName: "Sofia",
		lastName:  "I.",
		country:   "Argentina",
		continent: "Americas",
		age:       21,
		language:  "Java"},
	{
		firstName: "Lukas",
		lastName:  "X.",
		country:   "Croatia",
		continent: "Europe",
		age:       35,
		language:  "Ruby"},
	{
		firstName: "Madison",
		lastName:  "U.",
		country:   "United States",
		continent: "Americas",
		age:       42,
		language:  "Ruby"},
}

func main() {
	fmt.Println(getAverageAge(list1))
}

func getAverageAge(l []Developer) int {
	var sum int

	for _, dev := range l {
		sum += dev.age
	}

	return sum / len(l)
}
