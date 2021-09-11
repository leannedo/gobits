/* FIND ADMIN OF THE LANGUAGE CLAN */
package main

import "fmt"

type Developer struct {
	firstName, lastName, country, continent, language string
	age                                               int
	isAdmin                                           bool
}

var list1 = []Developer{
	{
		firstName: "Sofia",
		lastName:  "I.",
		country:   "Argentina",
		continent: "Americas",
		age:       21,
		language:  "Java",
		isAdmin:   true},
	{
		firstName: "Lukas",
		lastName:  "X.",
		country:   "Croatia",
		continent: "Europe",
		age:       35,
		language:  "Ruby",
		isAdmin:   false},
	{
		firstName: "Madison",
		lastName:  "U.",
		country:   "United States",
		continent: "Americas",
		age:       42,
		language:  "Go",
		isAdmin:   true},
}

func main() {
	fmt.Println(findAdmin(list1, "Kotlin"))
	fmt.Println(findAdmin(list1, "Go"))
}

func findAdmin(list []Developer, lang string) []Developer {
	var res []Developer
	for _, dev := range list {
		if dev.isAdmin && dev.language == lang {
			res = append(res, dev)
		}
	}

	return res
}
