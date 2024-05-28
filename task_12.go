package main

import "fmt"

func task_12() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	set := []string{}
	for _, val := range data {
		isUniq := true
		for _, i := range set {
			if val == i {
				isUniq = false
			}
		}
		if isUniq {
			set = append(set, val)
		}
	}
	fmt.Println(set)
}
