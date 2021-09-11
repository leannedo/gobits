package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Join([]string{"hello", "Finland", "sunny"}, " - ")
	fmt.Println(s)
}
