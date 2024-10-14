package main

import "fmt"

func createIntersection(inStrings []string) {

	var stringsMap = make(map[string]int)
	for _, elem := range inStrings {
		stringsMap[elem] += 1
	}

	fmt.Println(stringsMap)
}

func main() {
	inStrings := []string{"cat", "cat", "dog", "cat", "tree"}
	createIntersection(inStrings)
}
