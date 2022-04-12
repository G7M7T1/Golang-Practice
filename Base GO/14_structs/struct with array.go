package main

import "fmt"

type person5 struct {
	name string
	age  int
	like []string
}

func main() {
	p1 := person5{
		name: "John",
		age:  25,
		like: []string{"coding", "running"},
	}
	fmt.Println(p1)

	for i, v := range p1.like {
		fmt.Println(i, v)
	}
}
