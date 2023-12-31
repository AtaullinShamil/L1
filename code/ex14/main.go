package main

import "fmt"

func checkType(param interface{}) {
	// Проверяем тип
	switch v := param.(type) {
	// Используем свитч для сравнения значений
	case int:
		fmt.Println(v, " is int")
	case string:
		fmt.Println(v, " is string")
	case bool:
		fmt.Println(v, " is bool")
	case chan bool:
		fmt.Println(v, " is chan")
	default:
		fmt.Println(v, " is unknown type")
	}
}

func main() {
	checkType(42)
	checkType("Hello, World!")
	checkType(true)
	ch := make(chan bool)
	checkType(ch)
}
