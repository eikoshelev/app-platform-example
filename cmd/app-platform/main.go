package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "hello from app")
		if err != nil {
			log.Fatalf("error printing message to response writer %s", err.Error())
			return
		}
	})

	log.Printf("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
