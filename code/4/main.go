package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Указываем количество воркеров через флаг -num
	var num int
	flag.IntVar(&num, "num", 5, "number of workers")
	flag.Parse()

	// Создаем веитгруппу и канал
	var wg sync.WaitGroup
	numChan := make(chan int)

	// В цикле запускаем воркеров, которые работают, пока канал открыт
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for job := range numChan {
				fmt.Println(i, " worker - ", job)
			}
			wg.Done()
		}(i)
	}

	for i := 0; i < 1000; i++ {
		numChan <- i
	}
	close(numChan)
	wg.Wait()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
