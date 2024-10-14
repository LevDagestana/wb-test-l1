package main

import (
	"fmt"
	"sync"
)

func MulNum(num int, mulX chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	mulX <- num * 2
}

func read(mulX chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case num, ok := <-mulX:
			if !ok {
				return
			}
			fmt.Println(num)
		default:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var x = make(chan int, 10)
	var mulX = make(chan int, 10)

	for _, num := range nums {
		x <- num
	}
	close(x)

	for num := range x {
		wg.Add(1)
		go MulNum(num, mulX, &wg)
	}

	wg.Add(1)
	go read(mulX, &wg)

	wg.Wait()
	close(mulX)
}
