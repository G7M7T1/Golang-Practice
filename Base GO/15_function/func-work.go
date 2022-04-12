package main

import "fmt"

func main() {
	z := add(1, 2)
	fmt.Println(z)

	sayHello("Hi")

	x, y := moused("Hello", "World")
	fmt.Println(x)
	fmt.Println(y)
}

func add(x int, y int) int {
	return x + y
}

func sayHello(s string) {
	fmt.Println("Hello", s)
}

func moused(fn string, ln string) (string, bool) {
	aa := fmt.Sprint("Hello ", fn, ln)
	bb := true
	return aa, bb
}
