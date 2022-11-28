// printing ascii/UTF-8 encoded characters

package main

import (
	"fmt"
)

func main() {
	for i := 5000; i <= 5100; i++ {
		fmt.Printf("%d\t%x\t%#U\n", i, i, i)
	}
}
