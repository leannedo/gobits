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
		age:       35,
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
		age:       32,
		language:  "Ruby"},
}

func main() {
	fmt.Println(getFirstPython(list1))
}

func getFirstPython(list []Developer) string {
	result := "There is no Python developer"

	for _, dev := range list {
		if dev.language == "Python" {
			result = fmt.Sprintf("%s %s", dev.firstName, dev.lastName)
		}
	}

	return result
}
