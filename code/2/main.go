package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	// Создаем веитгруппу
	var wg sync.WaitGroup

	// В цикле запускаем горутины
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			result := n * n
			fmt.Println(result)
			wg.Done()
		}(num)
	}
	// Ожидаем выполнения горутин
	wg.Wait()
}
