package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			myMap[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(myMap)
}
