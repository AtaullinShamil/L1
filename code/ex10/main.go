package main

import "fmt"

func main() {
	// Создаем мапу для хранения данных
	myMap := make(map[int][]float64)

	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Проходим циклом по данным
	for _, val := range nums {
		// Получаем целое значение ключа группы
		group := int(val) / 10 * 10
		// Добавляем значение в группу
		myMap[group] = append(myMap[group], val)
	}
	fmt.Println(myMap)
}
