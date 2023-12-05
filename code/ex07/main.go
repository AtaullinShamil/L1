package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мапу
	myMap := make(map[int]int)
	// Создаем мьютекс
	mu := sync.Mutex{}
	// Веитгрупппа для корректного закрытия горутин
	wg := sync.WaitGroup{}

	// В цикле создаем 5 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			// Блокируем мьютекс
			mu.Lock()
			// Выполняем запись
			myMap[i] = i
			// Открываем мьютекс
			mu.Unlock()
			wg.Done()
		}(i)
		// Данные отправляем напрямую, чтобы не было замыкания
	}
	wg.Wait()
	fmt.Println(myMap)
}
