package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func main() {
	http.HandleFunc("/ping", handler)
	log.Print("App Platform test app")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "response count: %d", counter)
	log.Printf("count: %d", counter)
	counter++
}
