// FUNCTIONS continued
// CALLBACKS - passing a function as an argument/defining a function as a parameter in another function

package main

import (
	"fmt"
)

func main() {
	fmt.Println("heyo.")
	some_sum := sum(even, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(some_sum)
}

func sum(my_func func(my_nums ...int) []int, nums ...int) int {
	// demonstrate use of callbacks
	var my_sum int
	even_nums := my_func(nums...) // even expects a bunch of ints, not a slice of int

	for _, value := range even_nums {
		my_sum += value
	}

	return my_sum
}

func even(my_nums ...int) []int {
	my_even_nums := []int{}

	for _, value := range my_nums {
		if value%2 == 0 {
			my_even_nums = append(my_even_nums, value)
		}
	}
	return my_even_nums
}

// CLOSURES - used for limiting scopes of variables
