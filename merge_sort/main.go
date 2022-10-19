package main

import (
	"fmt"
)

func mergeSort(items []int) []int {
	if len(items) >= 2 {
		left := mergeSort(items[:len(items)/2])
		right := mergeSort(items[len(items)/2:])
		return arrangeThem(left, right)
	}
	return items
}

func arrangeThem(left []int, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}
	return result
}

func main() {
	originalData := []int{4, 7, 0, 2, 8, 5, 6, 3, 9, 1}
	fmt.Printf("Original Data: %v\n", originalData)

	sortedData := mergeSort(originalData)
	fmt.Printf("Sorted Data: %v\n", sortedData)
	// sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
}
