package main

import "fmt"

func task_10() {
	temp := make(map[int][]float64)
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, value := range data {
		key := int(value/10) * 10
		temp[key] = append(temp[key], value)
	}
	fmt.Println(temp)
}
