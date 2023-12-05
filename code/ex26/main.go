package main

import "fmt"

func Unique(str []rune) bool {
	// Мапа для хранения количества повторений символов
	myMap := make(map[rune]int)

	// Заполняем мапу
	for _, val := range str {
		myMap[val]++
	}

	// Проверяем количество повторений
	for _, value := range myMap {
		if value != 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Unique([]rune("abcd")))
	fmt.Println(Unique([]rune("abCdefAaf")))
	fmt.Println(Unique([]rune("aabcd")))
	fmt.Println(Unique([]rune("абвгд")))
	fmt.Println(Unique([]rune("аабвгд")))
}
