// maps - like python dictionaries - underlying structure is a hash table
// maps are unordered
// maps are a reference type i.e. think pointers

// syntax: map[string]int where string is the type of the key and
// int is the type of the value in the key-value pair

package main

import (
	"fmt"
)

func main() {
	// declaring a map variable, don't do this though as it returns nil -> no elements can be added
	// var my_map map[string]int

	// alternate declaration uses the shorthand operator ':=' omitting 'var' statement
	// use this instead
	my_map := make(map[string]int)

	// initialising values to the created map
	my_map["kpi 1 score"] = 50
	my_map["kpi 2 score"] = 57

	// printing out what is storede in the map
	for key, value := range my_map {
		fmt.Println(key, ":", value)
	}

	// deleting a key-value pair from the map (works even if the key is already absent)
	delete(my_map, "kpi 2 score") // 2 arguments, the map and the key

	for key, value := range my_map {
		fmt.Println(key, ":", value)
	}

	// checking if a certain key exists in the map
	_, ok := my_map["kpi score 3"] // blank identifier throws away first return i.e. the value matching the passed key - 'comma ok idiom'
	fmt.Println(ok)                // returns a boolean

	// composite literal method of initialising a map
	my_other_map := map[string]int{
		"kpi 3 score": 60,
		"kpi 4 score": 65,
		"kpi 5 score": 70,
	}

	for key, value := range my_other_map {
		fmt.Println(key, ":", value)
	}
}
