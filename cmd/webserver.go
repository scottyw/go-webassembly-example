package main

import (
	"fmt"
	"log"
	"net/http"
)

// Serve the contents of the "static" dir
func main() {
	fmt.Println("Starting webserver on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("static")))
	log.Fatalln(err)
}
