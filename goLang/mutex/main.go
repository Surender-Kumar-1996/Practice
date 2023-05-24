package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t: ", runtime.NumCPU())
	fmt.Println("Goroutines\t: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
		}()
		fmt.Println("Goroutines\t: ", runtime.NumGoroutine())
	}
	fmt.Println("Counter\t: ", counter)
	fmt.Println("Goroutines\t: ", runtime.NumGoroutine())
	wg.Wait()
}
