// don't communicate by sharing, share by communicating
// values are passed around channels and goroutines, and
// only one goroutine has access to a variable's value at any given time

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// initialising an identifier of type WaitGroup, see below
var wg sync.WaitGroup

func main() {
	print_stats()
	/* to launch sth into its own goroutine,
	precede the function call with keyword 'go' */
	wg.Add(1) // committing one goroutine to WaitGroup wg
	go foo()
	/* since Go is so fast, it executes the code below
	and func main exits before the above goroutine
	is even launched */
	bar()
	print_stats()
	/* one way to ensure the program 'waits' for the goroutine to execute
	before shutting down the program is to use 'Wait'Groups:
	wg.Add() above, wg.Wait() below, wg.Done() in foo()
	-> see docs, sync package */
	wg.Wait()
	/* when all the goroutines in the WaitGroups run wg.Done,
	the program stops waiting and proceeds execution */
}

func print_stats() {
	fmt.Println("OS:\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func foo() {
	for i := 1; i < 11; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // releases this goroutine from the WaitGroup
}

func bar() {
	for i := 1; i < 11; i++ {
		fmt.Println("bar:", i)
	}
}
