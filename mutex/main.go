package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	// to launch 100 goroutines
	const gortns = 100
	var wg sync.WaitGroup
	wg.Add(gortns)

	/*
		each goroutine is meant to increment the value of counter by 1,
		however, multiple concurrent goroutines trying to access the same variable
		leads to DATA RACES (see graphic)
		in order to avoid data races, mutices (from mutex - mutual exclusion i think)
		* like checking out a book in the library *
	*/
	var my_mutex sync.Mutex

	for i := 0; i < gortns; i++ {
		go func() {
			// locks this code block, which has access to counter variable
			my_mutex.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			// unlocks it
			my_mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
