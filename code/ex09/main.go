package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем два канала
	firstChan := make(chan int)
	secondChan := make(chan int)

	// Создаем две веитгруппы
	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Горутина, которая принимает данные из первого канала и отправляет во второй
	wg1.Add(1)
	go func() {
		for val := range firstChan {
			secondChan <- val * 2
		}
		wg1.Done()
	}()

	// Горутина, которая принимает данные из второго канала и печатает их
	wg2.Add(1)
	go func() {
		for val := range secondChan {
			fmt.Println(val)
		}
		wg2.Done()
	}()

	// Отправляем данные в канал
	for _, val := range nums {
		firstChan <- val
	}
	// Закрываем канал
	close(firstChan)
	// Ожидаем завершения первой горутины
	wg1.Wait()
	// Закрываем второй канал
	close(secondChan)
	// Ожидаем завершения второй горутины
	wg2.Wait()

}
