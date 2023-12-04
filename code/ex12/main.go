package main

import "fmt"

func main() {
	// Список строк
	list := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества
	set := make(map[string]bool)
	for _, val := range list {
		set[val] = true
	}

	// Вывод множества
	for item := range set {
		fmt.Println(item)
	}
}
