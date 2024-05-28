package main

import (
	"fmt"
	"strings"
)

func task_20() {
	str := "snow dog sun"
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	fmt.Println(words)
}
