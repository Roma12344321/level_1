package main

import "fmt"

func task_13() {
	a, b := 1, 2
	fmt.Printf("a=%d b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d b=%d\n", a, b)
}
