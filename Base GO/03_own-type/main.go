package main

import "fmt"

func main() {
	type NewType int
	var b NewType = 3
	fmt.Println(b)
	fmt.Printf("Type: %T\n", b)
}
