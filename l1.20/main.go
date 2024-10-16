package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	stringSlice := strings.Split(s, " ")
	var result string

	for i := len(stringSlice) - 1; i >= 0; i-- {
		elem := stringSlice[i]
		result += elem + " "
	}
	return result
}

func main() {
	inString := "snow dog sun"

	fmt.Println(reverseString(inString))
}
