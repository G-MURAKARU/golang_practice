// channels -> think relay runners and passing batons
// CHANNELS BLOCK: receiver and sender have to be in sync

package main

import (
	"fmt"
)

func main() {
	// code below does not work - DEADLOCK
	my_channel := make(chan int, 1) // buffer channel created with a capacity of one

	my_channel <- 42
	my_channel <- 43

	fmt.Println(<-my_channel)
}
