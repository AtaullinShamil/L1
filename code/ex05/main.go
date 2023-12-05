package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Время, которое будет выполняться программа
	var n int
	flag.IntVar(&n, "n", 2, "working time")
	flag.Parse()

	// Таймер выполнения программы
	timer := time.NewTimer(time.Duration(n) * time.Second)

	//канал для передачи данных
	numChan := make(chan int)

	// Веитгруппа для закрытия горутины
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Пока канал открыт, данные получаются
		for val := range numChan {
			fmt.Println(val)
		}
		wg.Done()
	}()

	for {
		select {
		// При срабатывании таймера, горутины завершаются, программа закрывается
		case <-timer.C:
			close(numChan)
			wg.Wait()
			return
			// Данные отправляются в канал
		default:
			numChan <- rand.Intn(100)
		}
	}
}
