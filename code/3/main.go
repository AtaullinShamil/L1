package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	// Создаем веитгруппу
	var wg sync.WaitGroup
	//Создаем канал для передачи чисел
	numChan := make(chan int)

	// В цикле запускаем горутины и отправляем результат в канал
	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			result := num * num
			numChan <- result
			wg.Done()
		}(num)
	}

	// В отдельной горутине ожидаем выполнения всех горутин для закрытия канала
	go func() {
		wg.Wait()
		close(numChan)
	}()

	// В цикле суммируем результаты, полученные из канала
	var sum int
	for numFromChan := range numChan {
		sum += numFromChan
	}
	fmt.Println(sum)
}
