package main

import "fmt"

// Bubblesort Algorithm
func bubbleSort(numbers []int) []int {
	if numbers != nil && len(numbers) > 0 {
		for i := 0; i < len(numbers); i++ {
			for j := 0; j < len(numbers); j++ {
				if numbers[i] < numbers[j] {
					temp := numbers[i]
					numbers[i] = numbers[j]
					numbers[j] = temp
				}
			}
		}
	} else {
		fmt.Println("Empty Array")
	}
	return numbers
}

// main function
func main() {
	//input: 4, 2, 1, 5, 3
	numbers := []int{4, 2, 1, 5, 3}

	//output: 1, 2, 3, 4, 5
	fmt.Println(bubbleSort(numbers))
}
