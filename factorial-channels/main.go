package main

import "fmt"

func main() {
	my_factorial := factorial(4)

	// this blocks us from exiting main before goroutines are done
	for element := range my_factorial {
		fmt.Println(element)
	}
}

// factorial finds the nth factorial and loads it onto an int channel 'output'
func factorial(n int) chan int {
	output := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		output <- total
		close(output)
	}()
	return output
}
