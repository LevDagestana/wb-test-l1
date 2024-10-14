package main

import (
	"fmt"
	"strconv"
)

func checkBit(num int64, pos int64) bool {
	v := num & (1 << pos)

	return v > 0
}

func setBit(num int64, pos int64) int64 {

	if checkBit(num, pos) == true {
		num &^= 1 << pos
	} else {
		num |= 1 << pos
	}
	return num
}

func main() {

	var num int64 = 40

	var position int64 = 1

	fmt.Println("Исходная последовательность битов", strconv.FormatInt(num, 2))

	changedNum := setBit(num, position)

	fmt.Println("Последовательность битов после функции setBit", strconv.FormatInt(changedNum, 2))
}
