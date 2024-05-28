package main

import (
	"fmt"
)

func task_8() {
	fmt.Printf("Введите число: ")
	var num int64
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Printf("Введите позицию бита: ")
	var bit int
	_, err = fmt.Scan(&bit)
	if err != nil {
		return
	}
	fmt.Printf("Введите значение бита (0 или 1): ")
	var val int
	_, err = fmt.Scan(&val)
	if err != nil {
		return
	}
	if val == 0 {
		num = num &^ (1 << bit)
	} else {
		num = num | (1 << bit)
	}
	fmt.Printf("Результат: %d (в двоичном виде: %b)\n", num, num)
}
