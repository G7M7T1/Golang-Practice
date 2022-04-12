package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is The Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is The About Page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0
	f, err := divideValues(float32(x), float32(y))

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can't not divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Server are running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
