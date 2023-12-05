package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	// Делим строку на слова
	chopped := strings.Split(str, " ")
	result := ""

	// Проходимся в обратном порядке и добавляем слова в результат
	for i := len(chopped) - 1; i >= 0; i-- {
		result = result + chopped[i]
		if i != 0 {
			result = result + " "
		}
	}
	fmt.Println(result)
}
