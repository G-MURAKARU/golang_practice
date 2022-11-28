// SORTING slices with package 'sort'

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type ByAge []Person

// ByAge implements the sort.Interface for []Person based on age field
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	people := []Person{
		{
			name: "Gicheru",
			age:  24,
		},
		{
			name: "Maina",
			age:  17,
		},
		{
			name: "Sylvia",
			age:  50,
		},
		{
			name: "Waire",
			age:  28,
		},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].age > people[j].age
	})
	fmt.Println(people)
}
