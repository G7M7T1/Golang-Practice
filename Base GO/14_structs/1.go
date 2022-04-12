package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
}
