package main

import (
	"fmt"
	"sort"
)

type developer struct {
	firstName, lastName, language string
}

type developerList []developer

func (d developerList) Len() int {
	return len(d)
}

func (d developerList) Less(i, j int) bool {
	// sort alphabetically
	if d[i].language == d[j].language {
		return d[i].firstName < d[j].firstName
	}
	return d[i].language < d[j].language
}

func (d developerList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	dList := developerList{
		{
			firstName: "Sofia",
			lastName:  "I.",
			language:  "Java"},
		{
			firstName: "Lukas",
			lastName:  "X.",
			language:  "Ruby"},
		{
			firstName: "Madison",
			lastName:  "U.",
			language:  "Go"},
		{
			firstName: "Austin",
			lastName:  "O.",
			language:  "Go"},
		{
			firstName: "Ngoc",
			lastName:  "D.",
			language:  "Kotlin"},
	}

	fmt.Println(sortByLanguage(dList))
}

func sortByLanguage(list developerList) developerList {
	sort.Sort(developerList(list))

	return list
}
