package main

import (
	"fmt"
)

func main() {
	var arr = []int{9, 5, 1, 30, 12}

	// arr = bubbleSort(arr)
	arr = quickSort(arr)
	fmt.Println(arr)
}

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := []int{}
	left := []int{}
	right := []int{}

	flag := arr[0]
	mid = append(mid, flag)
	for _, v := range arr[1:] {
		if v < flag {
			left = append(left, v)
		} else if v > flag {
			right = append(right, v)
		} else {
			mid = append(mid, v)
		}
	}
	left, right = quickSort(left), quickSort(right)
	res := append(append(left, mid...), right...)
	return res
}

// 冒泡排序
func bubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
