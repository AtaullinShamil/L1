package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before: a =", a, ", b =", b)
	// Меняем местами без создания дополнительной переменной
	a, b = b, a
	fmt.Println("After: a =", a, ", b =", b)
}
