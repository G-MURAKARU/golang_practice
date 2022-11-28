// FUNCTIONS - THEY ARE FIRST-CLASS CITIZENS
// METHODS
// INTERFACES
// ANONYMOUS FUNCTIONS - SELF-EXECUTING FUNCTIONS
// FUNC EXPRESSIONS
// DEFERRING A FUNCTION
// SLICE UNFURLING
// VARIADIC FUNCTIONS

package main

import (
	"fmt"
)

// syntax: func (receiver) identifier(parameters) (return_type(s)) {
//				code
// 			}

// variadic parameters: ...

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

// INTERFACES
// interfaces allow us to define behaviour - polymorphism
// empty interface - every identifier has no methods - everything implements the type empty-interface

// e.g. any type that has access to method 'speak()' is also of type 'human'
// since both Person and Double_Zero have the method speak(),
// identifiers of their types are also of type human
type human interface {
	// polymorphism - a value can be of more than one type
	speak()
}

func main() {
	agent1 := Double_Zero{
		// the unqualified type name acts as the field name
		Person: Person{
			first: "Rick",
			last:  "Sanchez",
		},
		target:          "Birdman",
		license_to_kill: true,
	}

	person1 := Person{
		// if any fields are not initialised, they will default to the zero value of their type
		first: "Gicheru",
		last:  "Murakaru",
	}

	// deferring a function (see golang spec)
	defer bar()
	defer agent1.speak()
	defer person1.speak() // polymorphism in the speak() method
	// notice defer's LIFO behaviour - person1.speak() runs before agent1.speak() which runs before bar()

	// ANONYMOUS FUNCTION - CLOSURES
	// parameters and return are optional
	truth := func(x int) bool {
		fmt.Println("anonymous function ran")
		if y := x == 42; y {
			return true
		} else {
			return false
		}
	}(43) // SELF-EXECUTION

	if truth {
		fmt.Println("truth")
	} else {
		fmt.Println("no truth")
	}

	// FUNC EXPRESSION - assigning a function to a variable
	f := func() {
		fmt.Println("func expression ran")
	}
	// calling the function
	f()

	full, add := foo("Gicheru", "Murakaru", 1, 2, 3, 4, 5)
	fmt.Printf("%s\n%d\n", full, add)

	// interfaces
	sth(agent1)
	sth(person1)
}

func foo(fn string, ln string, num ...int) (string, int) {
	// num comes in as a slice of int
	full_name := fmt.Sprint(fn, " ", ln)

	var sum int

	for _, value := range num {
		sum += value
	}

	return full_name, sum
}

// DEFER KEYWORD
func bar() {
	fmt.Println("bar")
}

// METHODS
func (agent Double_Zero) speak() {
	// the receiver above attaches the function speak() to identifiers of type Double_Zero
	// think object methods
	fmt.Printf("I am %s %s.\n", agent.first, agent.last)
	fmt.Println("You are excommunicado.")
}

func (mtu Person) speak() {
	// the receiver above attaches the function speak() to identifiers of type Double_Zero
	// think object methods
	fmt.Printf("I am %s %s.\n", mtu.first, mtu.last)
	fmt.Println("Hey.")
}

// interfaces
func sth(h human) {
	// type assertion - since types Person and Double_Zero are also type human, helps to distinguish
	switch h.(type) { // you can also switch on type, not just bool
	case Person:
		fmt.Printf("I'm a person. My name is %s.\n", h.(Person).first)
	case Double_Zero:
		fmt.Printf("I'm a secret agent. My name is %s.\n", h.(Double_Zero).first)
	}
}
