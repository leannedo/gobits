/* COUNT THE NUMBERS OF DEVELOPERS COMING FROM EUROPE AND CODE IN GO */
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
		language:  "Go"},
	{
		firstName: "Lukas",
		lastName:  "X.",
		country:   "Croatia",
		continent: "Europe",
		age:       35,
		language:  "Go"},
	{
		firstName: "Madison",
		lastName:  "U.",
		country:   "United States",
		continent: "Americas",
		age:       42,
		language:  "Ruby"},
}

func main() {
	fmt.Println(countDevelopers(list1))
}

func countDevelopers(l []Developer) int {
	sum := 0
	for _, dev := range l {
		if dev.language == "Go" && dev.continent == "Europe" {
			sum++
		}
	}

	return sum
}
