package main

import "fmt"

func binarySearch(arr []int, lookingFor int) (int, bool) {
	min := 0
	max := len(arr) - 1
	for min <= max {
		mid := (min + max) / 2
		if lookingFor == arr[mid] {
			// Если число найдено, возвращаем его порядковый номер и true
			return mid, true
		} else {
			// Если число меньше половины, уменьшаем максимум
			if arr[mid] > lookingFor {
				max = mid - 1
			} else {
				// Если число больше половины, увеличиваем минимум
				min = mid + 1
			}

		}
	}
	// Если число отсутствует, возвращаем false
	return 0, false
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(nums, 2))
}
