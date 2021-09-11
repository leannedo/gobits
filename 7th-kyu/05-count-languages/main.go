package main

import (
	"fmt"
)

type developer struct {
	firstName, lastName, country, continent, language string
	age                                               int
}

var list1 = []developer{
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
	fmt.Println(Count(list1))
}
