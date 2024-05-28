package main

import (
	"fmt"
	"math/big"
)

func task_22() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("10485772", 10)
	b.SetString("10485211", 10)
	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сложение: %s\n", sum.String())
	sub.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", sub.String())
	mul.Mul(a, b)
	fmt.Printf("Умножение: %s\n", mul.String())
	div.Quo(a, b)
	fmt.Printf("Деление: %s\n", div.String())
}
