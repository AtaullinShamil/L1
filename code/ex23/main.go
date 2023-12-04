package main

import "fmt"

func main() {
	// Объявляем слайс
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// i-тый элемент
	i := 2

	// Удаляем i-тый элемент
	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)
}
