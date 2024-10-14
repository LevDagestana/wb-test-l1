package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	num1 := 15
	num2 := 50
	fmt.Println("До swap:")
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	num1, num2 = swap(num1, num2)

	fmt.Println("После swap:")
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
}
