package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем большие числа
	a := big.NewInt(1048577)
	b := big.NewInt(2)

	// Выполняем арифметические операции
	var result big.Int

	fmt.Println("a + b =", result.Add(a, b))

	fmt.Println("a - b =", result.Sub(a, b))

	fmt.Println("a * b =", result.Mul(a, b))

	fmt.Println("a / b =", result.Div(a, b))
}
