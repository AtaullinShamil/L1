package main

import (
	"flag"
	"fmt"
	"math/rand"
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
			// Цикл существует, пока открыт канал
			for job := range numChan {
				fmt.Println(i, " worker - ", job)
			}
			fmt.Println(i, " worker stopped")
			wg.Done()
		}(i)
	}

	// Перехватываем сигнал control+c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for {
		select {
		// Если сигнал получен, закрываем канал и завершаем горутины
		case <-sigs:
			{
				close(numChan)
				wg.Wait()
				return
			}
		// Если сигнала нет, отправляем данные в канал
		default:
			numChan <- rand.Intn(100)
		}
	}
}
