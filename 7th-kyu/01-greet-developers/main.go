package main

import "fmt"

type Developer struct {
	firstName, lastName, country, continent, language string
	age                                               int
}

type developerWithGreeting struct {
	Developer
	greeting string
}

var list1 = []Developer{
	{
		firstName: "Sofia",
		lastName:  "I.",
		country:   "Argentina",
		continent: "Americas",
		age:       35,
		language:  "Java"},
	{
		firstName: "Lukas",
		lastName:  "X.",
		country:   "Croatia",
		continent: "Europe",
		age:       35,
		language:  "Python"},
	{
		firstName: "Madison",
		lastName:  "U.",
		country:   "United States",
		continent: "Americas",
		age:       32,
		language:  "Ruby"},
}

func main() {
	fmt.Println(greetDevelopers(list1))
}

func greetDevelopers(list []Developer) []developerWithGreeting {
	var listWithGreeting []developerWithGreeting

	for _, dev := range list {
		message :=
			fmt.Sprintf("Hello %s, what do you think about %s ?", dev.firstName, dev.language)

		newD := developerWithGreeting{
			Developer: dev,
			greeting:  message,
		}
		listWithGreeting = append(listWithGreeting, newD)
	}

	return listWithGreeting
}
