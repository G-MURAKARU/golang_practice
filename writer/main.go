// THE WRITER INTERFACE - deep dive into using the standard library
// interfaces - anything that implements the Write method of the Writer interface
// is ALSO OF TYPE WRITER

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("heyo.") // syntax: func Println(a ...interface{}) { return Fprintln(os.Stdout, a...) }

	// Stdout has a NewFile() somewhere and NewFile returns a *File (pointer to a File)
	// and anything of type *File implements the Write function of the Writer interface (see docs)
	// meaning anything of type *File is also of type Writer
	// therefore, os.Stdout is also of type Writer

	fmt.Fprintln(os.Stdout, "heyo.")   // syntax: func Fprintln(w Writer, a...interface{}) (n int, err error)
	io.WriteString(os.Stdout, "heyo.") // syntax: func WriteString(w Writer, s string) (n int, err error)
}
