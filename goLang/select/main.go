package main

import (
	"fmt"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan struct{})

	go send(odd, even, quit)

	fmt.Println("Receiving form the channel")

	func(o, e <-chan int, q <-chan struct{}) {
		for {
			select {
			case v := <-e:
				fmt.Println("Receiving from eve channel: ", v)
			case v := <-o:
				fmt.Println("Receiving fomr odd channel: ", v)
			case i, ok := <-q:
				fmt.Println("Receiving from quit channel", i, ok)
				return
			}
		}
	}(even, odd, quit)

	fmt.Println("About to exit.....")
}

func send(odd, even chan<- int, quit chan<- struct{}) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
			// go func(even chan<- int, i int) {//use waitgroups to capture the output of all the go routines
			// 	even <- i
			// }(even, i)
		} else {
			odd <- i
			// go func(odd chan<- int, i int) { //use waitgroups to capture the output of all the go routines
			// 	odd <- i
			// }(odd, i)
		}
	}
	quit <- struct{}{}
}
