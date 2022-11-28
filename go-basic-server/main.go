package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello_handler(writer http.ResponseWriter, request *http.Request) {
	// if the request path is not to the /hello route, page not found...
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}
	// only GET requests are allowed on the /hello route
	if request.Method != "GET" {
		http.Error(writer, "Method Is Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Hello!")
}

func form_handler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(writer, "Post request successful.\n")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "Name: %s\n", name)
	fmt.Fprintf(writer, "Address: %s\n", address)
}

func main() {
	//points golang to the static directory to find the index.html file
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/form", form_handler)
	http.HandleFunc("/hello", hello_handler)

	fmt.Println("Starting server at port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
