package main

import (
	"fmt"
)

// defaults to the zero value of its type
// i.e. false for bool, zero for int, "" for string, etc
var my_bool bool

// in Go, you can create your own types, e.g.
type gicheru int

func main() {
	var my_string string
	my_string = "Hello"
	my_bool = true
	my_integer := 73

	my_bool = print_func(my_bool, my_integer, my_string)

	if !my_bool {
		my_other_string := "World"
		fmt.Println(my_other_string)
	} else {
		fmt.Println("failure.")
	}

	var awesome gicheru = 100
	fmt.Println(awesome)
}

func print_func(my_bool bool, my_int int, my_str string) bool {
	if my_bool {
		fmt.Println(my_int, my_str)
	}

	my_bool = false
	return my_bool
}
