package main

import (
	"net/http"
	"fmt"
	"log"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}

func main() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
