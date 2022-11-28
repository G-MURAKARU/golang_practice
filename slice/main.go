// slices (and arrays) - stores values of one type
// for an array, the size is part of its TYPE
// SLICES ARE BUILT ON TOP OF ARRAYS

package main

import (
	"fmt"
)

func main() {
	// composite literal
	// identifier = type{values}
	x := []int{1, 2, 3, 4, 5} // 'x' is a slice of ints 1, 2, 3, 4, 5
	fmt.Printf("%T\n", x)

	for index, value := range x {
		fmt.Println(index, value)
	}

	y := []int{6, 7, 8, 9, 10}

	x = append(x, y...) // y... unfurls slice y, like *args in Python

	for index, value := range x {
		fmt.Println(index, value)
	}

	// to delete from a slice, use append and extract slices from the parent slice then append them to each other
	x = append(x[0:2], x[4:]...) // deletes values at indices 2 and 3

	fmt.Println(x)

	// using the builtin function make
	// syntax: make([]T, len, cap)
	// returns an initialised (not zeroed) value of type T (not *T)
	// see Effective Go
	z := make([]int, 10, 100)
	fmt.Println(z)

	// multidimensional slices: xp := [][]string{[]string, []string}
}
