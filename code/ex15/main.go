package main

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

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) > 99 {
		justString = v[:100]
	} else {
		justString = v[:len(v)]
	}

}

func main() {
	someFunc()
}
