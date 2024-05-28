package main

import "fmt"

func task_19() {
	fmt.Println(reverseString("главрыба"))
}

func reverseString(str string) string {
	r := []rune(str)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}
