package main

import (
	"fmt"
	"sync"
)

func goSquare(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()

}
func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go goSquare(num, &wg)
	}

	wg.Wait()
}
