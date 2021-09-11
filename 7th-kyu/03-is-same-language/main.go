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

var list2 = []Developer{
	{
		firstName: "Sofia",
		lastName:  "I.",
		country:   "Argentina",
		continent: "Americas",
		age:       35,
		language:  "Ruby"},
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
	fmt.Println(isSameLanguage(list1))
	fmt.Println(isSameLanguage(list2))
}

func contain(s []string, language string) bool {
	for _, v := range s {
		if v == language {
			return true
		}
	}
	return false
}

func isSameLanguage(list []Developer) bool {
	var result []string

	for _, dev := range list {
		if !contain(result, dev.language) {
			result = append(result, dev.language)
		}
	}

	return len(result) == 1
}
