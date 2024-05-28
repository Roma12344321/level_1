package main

import (
	"fmt"
	"strings"
)

func task_15() {
	fmt.Printf("justString: %v\n", someFunc())
}

func createHugeString(size int) string {
	var builder strings.Builder
	builder.Grow(size)
	for i := 0; i < size; i++ {
		builder.WriteRune('t')
	}
	return builder.String()
}

func someFunc() string {
	v := createHugeString(1024)
	justString := string([]rune(v)[:100])
	return justString
}
