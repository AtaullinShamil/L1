package main

import "fmt"

func main() {
	// Множества
	nums1 := []int{0, 1, 2, 3, 4, 5}
	nums2 := []int{4, 5}

	// Мапы для проверки пересечения
	myMap1 := make(map[int]bool)
	myMap2 := make(map[int]bool)

	// Заполняем первую мапу
	for _, val := range nums1 {
		myMap1[val] = true
	}

	// Заполняем вторую мапу
	for _, val := range nums2 {
		myMap2[val] = true
	}

	// Создаем слайс для результата пересечений
	result := make([]int, 0)

	for key := range myMap1 {
		// Проверяем наличие пересечения
		if myMap2[key] {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
