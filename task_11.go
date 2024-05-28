package main

import "fmt"

func task_11() {
	set1 := []int{4, 23, 53, 2, 6, 9, 7}
	set2 := []int{6, 7, 3, 5, 45, 53}
	m := make(map[int]bool)
	var res []int
	for _, v := range set1 {
		m[v] = true
	}
	for _, v := range set2 {
		if m[v] == true {
			res = append(res, v)
		}
	}
	fmt.Println(res)
}
