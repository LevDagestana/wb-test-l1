package main

import (
	"fmt"
	"log"
	"time"
)

func chanReader(d chan string) {
	for {
		data, ok := <-d
		if !ok {
			fmt.Println("Канал закрыт, завершаю работу")
			return
		}
		fmt.Printf("Получил данные: %s\n", data)
	}
}

func main() {

	var sec int
	fmt.Println("Введите количество секунд: ")
	_, err := fmt.Scanln(&sec)

	if err != nil {
		log.Fatal(err)
	}

	d := make(chan string)

	go chanReader(d)

	go func() {
		for {
			var data string
			fmt.Println("Введите данные: ")
			fmt.Scanln(&data)
			d <- data
		}
	}()

	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println("Время работы вышло")

	close(d)
}
