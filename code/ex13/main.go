package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before: a =", a, ", b =", b)
	a, b = b, a
	fmt.Println("After: a =", a, ", b =", b)
}
