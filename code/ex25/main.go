package main

import (
	"fmt"
	"time"
)

// Функция возвращает канал, который закрывается после указанного времени и разблокирует горутину
func Sleep(tm time.Duration) {
	<-time.After(tm)
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("Finish")
}
