package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello, World!")
		if err != nil {
			log.Println(err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
