package main

import "fmt"

func task_17() {
	arr := []int{3, 5, 7, 8, 23, 54, 66, 78}
	k := 66
	fmt.Println(binarySearch(arr, k))
}

func binarySearch(arr []int, k int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		med := (left + right) / 2
		if arr[med] == k {
			return true
		}
		if arr[med] < k {
			left = med + 1
		} else {
			right = med - 1
		}
	}
	return false
}
