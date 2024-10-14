package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")
	Sleep(3)
	fmt.Println("Прошло 3 секунды")
}
