// channels -> think relay runners and passing batons
// CHANNELS BLOCK: receiver and sender have to be in sync

package main

func main() {
	/*
		SYNTAX:
		send & receive channel
		bi_channel := make(chan int)

		send only channel
		send_channel := make(chan<- int)

		receive only channel
		receive_channel := make(<-chan int)
	*/

	/*
		ASSIGNMENT
		general to specific ✅
		send_channel = bi_channel
		receive_channel = bi_channel
		specific to general ❎
		bi_channel = send_channel
		bi_channel = receive_channel
	*/
}
