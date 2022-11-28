// JSON marshalling - from struct to json

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// notice the field names are title case
	First  string
	Second string
	Age    int
}

// json marshal syntax:
// func json.Marshal(variable interface{}) ([]byte, error)
func main() {
	person1 := person{
		First:  "Gicheru",
		Second: "Murakaru",
		Age:    23,
	}

	person2 := person{
		First:  "Nyokabi",
		Second: "Waweru",
		Age:    23,
	}

	people := []person{person1, person2}

	json_data, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(json_data))
}
