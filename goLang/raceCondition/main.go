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

	for i := 0; i < gs; i++ {
		wg.Add(i)
		go func() {
			defer wg.Done()
			v := counter
			runtime.Gosched()
			v++
			counter = v
		}()
		fmt.Println("Goroutines\t: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter\t: ", counter)
	fmt.Println("Goroutines\t: ", runtime.NumGoroutine())
	wg.Done()
}
