package main

import "fmt"

type newman struct {
	name string
	age  int
}

func main() {
	p1 := newman{
		name: "James",
		age:  32,
	}
	fmt.Println(p1)
}
