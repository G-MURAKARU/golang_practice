/*
	RANGE WITH CHANNELS
	when looping over a channel with range, it is imperative that the channel is closed
	otherwise the program will crash, deadlocked channels
	because the range will continue until the channel is closed,
	even if there are no more values left on the channel (deadlock)
*/

package main

import (
	"fmt"
)

func main() {
	some_channel := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			some_channel <- i
		}
		// to close a channel
		close(some_channel)
	}()

	for v := range some_channel {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
