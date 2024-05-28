package main

import (
	"fmt"
)

func task_16() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println(quicksort(arr))
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivotIndex := len(arr) / 2
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}
