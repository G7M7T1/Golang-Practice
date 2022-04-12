package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number, err := fmt.Fprintf(w, "Hello World")

		if err != nil {
			log.Println("Err:", err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", number))
	})

	fmt.Println("Server are running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
