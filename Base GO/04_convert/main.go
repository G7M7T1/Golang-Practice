package main

import "fmt"

type NewType int

func main() {
	var a int
	var b NewType = 3

	a = 10
	fmt.Println(a)

	a = int(b)
	fmt.Println(a)

	fmt.Println(b)
	fmt.Printf("Type: %T\n", b)
}
