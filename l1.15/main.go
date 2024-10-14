package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(size int) string {
	return strings.Repeat("wb", size)
}

func main() {
	fmt.Println("justString:", someFunc())
}
