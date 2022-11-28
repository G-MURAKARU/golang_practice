// STRUCTS
// classic OOP: create classes, insantiate objects of that class
// Go OOP: create types, initialise identifiers of that type

// a struct is a sequence of named elements, called fields,
// each of which has a name and a type

// if a field is initialised without a name i.e. just a type, it is known
// as an embedded/anonymous type (used for inheritance)

// structs also have tags - working with json data

// to optimise performance, arrange fields in descending order from largest to smallest (memory footprint)
// e.g. int64 (8-byte) -> float32 (4-byte) -> bool (2-byte) -> int8 (1-byte)

package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
}

type Double_Zero struct {
	// the struct Person is an embedded struct in this struct Double_Zero
	Person
	target          string
	license_to_kill bool
}

func main() {
	person1 := Person{
		// if any fields are not initialised, they will default to the zero value of their type
		first: "Gicheru",
		last:  "Murakaru",
	}

	person2 := Person{
		first: "Shawn",
		last:  "Murakaru",
	}

	fmt.Println(person1)
	fmt.Println(person2.first)

	agent1 := Double_Zero{
		// the unqualified type name acts as the field name
		Person: Person{
			first: "Rick",
			last:  "Sanchez",
		},
		target:          "Birdman",
		license_to_kill: true,
	}

	fmt.Println(agent1)

	// to create an anonymous struct
	// if you want to use the struct in a small area
	// e.g. within a certain scope only
	// note: does not have a name
	villain := struct {
		Person  // embedded struct
		origin  string
		nemesis string
	}{
		// you can still use embedded structs (inheritance)
		Person: Person{
			first: "Loki",
			last:  "Odinson",
		},
		origin:  "Asgardian",
		nemesis: "Thor",
	}

	fmt.Println(villain)
}
