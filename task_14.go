package main

import (
	"fmt"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is chan int")
	case chan string:
		fmt.Println("Type is chan string")
	case chan bool:
		fmt.Println("Type is chan bool")
	default:
		fmt.Println("Unknown type")
	}
}

func task_14() {
	var i interface{} = 42
	determineType(i)
}
