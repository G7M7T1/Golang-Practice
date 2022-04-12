package main

import "fmt"

func main() {

	var a = 10

	switch {

	case a == 10:
		fmt.Println("This should print")
		fallthrough

	case a == 11:
		fmt.Println("This should not print")
		fallthrough

	case a == 12:
		fmt.Println("This should print")
		fallthrough

	case (3 == 4):
		fmt.Println("This should not print")
		fallthrough

	case (4 == 4):
		fmt.Println("This should print 4 4")

	default:
		fmt.Println("This is default case")
	}
}
