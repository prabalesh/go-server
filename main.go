package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, user!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("Server started at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
