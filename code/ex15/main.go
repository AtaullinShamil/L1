package main

import "fmt"

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Мы не знаем, строку какого размера возвращает createHugeString.
// Если она короче 100 символов, то произойдет ошибка
// Поэтому воспользуемся проверкой длинны

// Теоретически возможно переполнение

var justString string

// Реализация для примера
func createHugeString(size int) string {
	result := ""

	if size == 1024 {
		for i := 0; i < size; i++ {
			result = result + "!"
		}
	} else {
		result = "!"
	}
	return result
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) > 99 {
		justString = v[:100]
		fmt.Println(len(justString))
	} else {
		justString = v[:len(v)]
		fmt.Println(len(justString))
	}

}

func main() {
	someFunc()
}
