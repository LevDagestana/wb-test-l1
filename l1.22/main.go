package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(3)
	a.Exp(a, big.NewInt(20), nil)

	b := big.NewInt(4)
	b.Exp(b, big.NewInt(20), nil)

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", mul.String())

	div := new(big.Int).Div(mul, a)
	fmt.Printf("Деление: %s\n", div.String())

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", sum.String())

	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s\n", sub.String())
}
