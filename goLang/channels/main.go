package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//send
	go foo(ch)

	//receive
	bar(ch)

	fmt.Println("About to exit.....")
}

func foo(ch chan<- int) {
	ch <- 42
}

func bar(ch <-chan int) {
	fmt.Println("The value receivedd from channel is: ", <-ch)
}
