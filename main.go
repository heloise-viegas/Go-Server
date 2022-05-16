package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err)
		return
	}
	fmt.Fprintf(w, "POST successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name is %v and address is %v", name, address)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is starting at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
