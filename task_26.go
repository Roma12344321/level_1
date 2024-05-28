package main

import (
	"fmt"
	"strings"
)

func task_26() {
	set := make(map[rune]bool)
	str := "abCdefAaf"
	str = strings.ToLower(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		set[r[i]] = true
	}
	if len(r) == len(set) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
