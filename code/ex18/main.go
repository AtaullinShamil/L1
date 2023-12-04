package main

import (
	"fmt"
	"sync"
)

// Структура счетчика
type counter struct {
	times int
}

func main() {
	var wg sync.WaitGroup
	// Мьютекс для избежания data race
	var mu sync.Mutex

	// Инициализируем структуру
	var myCounter counter

	// Создаем горутины
	wg.Add(1)
	go func() {
		for i := 0; i < 500; i++ {
			// Блокируем мьютекс
			mu.Lock()
			// Увеличиваем значение
			myCounter.times++
			// Открываем мьютекс
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 500; i++ {
			mu.Lock()
			myCounter.times++
			mu.Unlock()
		}
		wg.Done()
	}()

	// Ожидаем завершения горутин
	wg.Wait()
	fmt.Println(myCounter.times)
}
