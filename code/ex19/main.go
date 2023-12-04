package main

import "fmt"

func main() {
	// Создаем слайс рун для поддержки unicode
	arr := []rune("главрыба")

	var result []rune

	// Проходимся по слайсу с конца
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	fmt.Println(string(result))

}
