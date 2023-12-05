package main

import "fmt"

// Функция которая в числе digit устанавливает бит с номером number на bitValue
func setBit(digit int, bitValue int, number int) int {
	if bitValue != 0 && bitValue != 1 {
		fmt.Println("bitValue can be only 1 or 0")
		return 0
	}

	// Маска в которой нужный бит установлен на 1
	mask := 1 << number
	if bitValue == 0 {
		// Если нужно установить 0, используем побитовое И НЕ
		digit = digit &^ mask
	} else {
		// Если нужно установить 1, используем побитовое ИЛИ
		digit = digit | mask
	}
	return digit
}

func main() {
	i := 0

	fmt.Printf("Before : %d - %10b\n", i, i)
	i = setBit(i, 1, 4)
	fmt.Printf("After 1 : %d - %10b\n", i, i)
	i = setBit(i, 1, 8)
	fmt.Printf("After 2 : %d - %10b\n", i, i)

}
