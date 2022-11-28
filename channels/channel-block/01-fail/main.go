// channels -> think relay runners and passing batons
// CHANNELS BLOCK: receiver and sender have to be in sync
// channels allow us to pass values between goroutines

package main

import (
	"fmt"
)

func main() {
	// code below does not work - DEADLOCK
	my_channel := make(chan int)

	my_channel <- 42

	fmt.Println(<-my_channel)
}
