// JSON UNMARSHAL - from json data to Go struct
// note - json encode and decode

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string `json:"First"`
	Second string `json:"Second"`
	Age    int    `json:"Age"`
}

// json unmarshal syntax
// func json.Unmarshal(data []byte, variable<pointer> interface{}) err
func main() {
	json_data := `[{"First":"Gicheru","Second":"Murakaru","Age":23},{"First":"Nyokabi","Second":"Waweru","Age":23}]`
	// json data taken in as a raw string literal

	json_bs := []byte(json_data)

	var people []person

	err := json.Unmarshal(json_bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(people)
	fmt.Printf("%T\n", people)
}
