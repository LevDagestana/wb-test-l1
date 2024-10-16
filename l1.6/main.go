package main

import (
	"context"
	"fmt"
	"time"
)

func cancelWithCtx(ctx context.Context) {

	for {
		select {

		default:
			fmt.Println("ctx: работаю...")

		case <-ctx.Done():
			fmt.Println("ctx: Завершаю работу")
			return
		}
	}

}

func cancelWithChan(stop chan bool) {

	for {
		select {

		default:
			fmt.Println("chan: работаю...")

		case <-stop:
			fmt.Println("chan: Завершаю работу")
			return
		}
	}

}

func main() {
	stop := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go cancelWithChan(stop)
	go cancelWithCtx(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel()
	stop <- true
}
