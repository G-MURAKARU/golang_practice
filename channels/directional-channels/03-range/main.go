package main

import (
	"fmt"
)

func main() {
	my_channel := make(chan int)

	// launches a goroutine
	go foo(my_channel)

	// note: not launched into another goroutine, but receives from the channel
	bar(my_channel)

	// due to program structure, no need for a WaitGroup

	fmt.Println("about to exit")
}

func foo(ch chan<- int) {
	// general to specific assignment
	ch <- 42
}

func bar(ch <-chan int) {
	// general to specific assignment

	/* code gets here and BLOCKS as it waits for foo() to run in its goroutine
	and "pass the channel baton" to bar once it is done */
	fmt.Println(<-ch)
}
