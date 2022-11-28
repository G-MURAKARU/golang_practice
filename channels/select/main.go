// SELECT STATEMENT (in func receive)
// - analogous to switch statement but for channels

package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit.")
}

func send(ev, od, qu chan<- int) {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			ev <- i
		} else {
			od <- i
		}
	}
	// close(ev)
	// close(od)
	qu <- 0
}

func receive(ev, od, qu <-chan int) {
	for { // infinite loop, broken out of when channels close
		select {
		case value := <-ev: // if value has been pulled off of the even channel
			fmt.Println("from even channel:", value)
		case value := <-od: // if value has been pulled off the odd channel
			fmt.Println("from odd channel:", value)
		case value := <-qu: // if value has been pulled off the quit channel
			fmt.Println("from quit channel:", value)
			return
		}
	}
}
