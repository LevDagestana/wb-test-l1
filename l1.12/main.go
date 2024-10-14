package main

import "fmt"

func findIntersection(dummyStrings []string) {

	var stringsMap = make(map[string]int)
	for _, elem := range dummyStrings {
		stringsMap[elem] += 1
	}

	fmt.Println(stringsMap)
}

func main() {
	dummyStrings := []string{"cat", "cat", "dog", "cat", "tree"}
	findIntersection(dummyStrings)
}
