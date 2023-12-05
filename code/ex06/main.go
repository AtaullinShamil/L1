package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Передача сигнала остановки
	stopChan := make(chan bool)
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("1 - stop signal")
				return
			default:
				fmt.Println("1 - working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	stopChan <- true
	time.Sleep(1 * time.Second)

	// Остановка с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("2 - stop ctx")
				return
			default:
				fmt.Println("2 - working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// Закрытие канала
	numChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _ = range numChan {
			fmt.Println("3 - working")
		}
		fmt.Println("3 - stop closing chan")
		wg.Done()
	}()
	for i := 1; i < 6; i++ {
		numChan <- i
		time.Sleep(time.Second)
	}
	close(numChan)
	wg.Wait()
}
