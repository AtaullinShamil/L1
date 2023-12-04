package main

import "fmt"

func quicksort(array []int) []int {
	// Базовый случай
	if len(array) < 2 {
		return array
	} else {
		// Рекурсивный случай

		// Опорный элемент
		pivot := array[0]
		// Меньше опорного элемента
		less := make([]int, 0)
		// Больше опорного элемента
		greater := make([]int, 0)
		for _, i := range array[1:] {
			if i <= pivot {
				less = append(less, i)
			} else {
				greater = append(greater, i)
			}
		}
		return append(append(quicksort(less), pivot), quicksort(greater)...)
	}
}

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 1}))
}
