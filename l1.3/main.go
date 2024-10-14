package main

import (
	"fmt"
	"sync"
)

func coSquare(num int) int {
	return num * num

}
func main() {
	numbers := []int{2, 4, 6, 8, 10.}
	var wg sync.WaitGroup
	var sum int
	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			sum += coSquare(num)
			wg.Done()
		}(num)

	}
	wg.Wait()
	fmt.Println(sum)
}
