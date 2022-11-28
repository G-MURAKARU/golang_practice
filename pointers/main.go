// POINTERS

// everything in Go is passed by value

package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("address of x:", &x)
	// passing the value x
	bar(x)
	fmt.Println("x passed with no pointer:", x)
	// passing the address where value 'x' is stored
	fmt.Println("---------------------")
	foo(&x)
	fmt.Println("x passed with pointer:", x)
}

func foo(y *int) {
	fmt.Println("foo:", *y, "address:", y)
	*y = 43
	fmt.Println("foo (assigned):", *y, "address:", y)
}

func bar(y int) {
	fmt.Println("bar:", y)
	y = 43
	fmt.Println("bar (assigned):", y)
}
