// channels -> think relay runners and passing batons

package main

import (
	"fmt"
)

func main() {
	// code below does not work - DEADLOCK
	my_channel := make(chan int)

	go func() {
		my_channel <- 42
	}()

	fmt.Println(<-my_channel)
}
