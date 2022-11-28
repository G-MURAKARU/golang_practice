// USING IOTA AND BIT SHIFTING

package main

import (
	"fmt"
)

const (
	// iota used for incrementing by 1
	_  = iota
	kb = 1 << (iota * 10) // kilobyte
	mb = 1 << (iota * 10) // megabyte
	gb = 1 << (iota * 10) // gigabyte
	tb = 1 << (iota * 10) // terabyte
)

func main() {
	fmt.Printf("%v\t\t\t%b\n", kb, kb)
	fmt.Printf("%v\t\t\t%b\n", mb, mb)
	fmt.Printf("%v\t\t\t%b\n", gb, gb)
	fmt.Printf("%v\t\t\t%b\n", tb, tb)
}
