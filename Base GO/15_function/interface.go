package main

import "fmt"

type Person4 struct {
	Name string
	Age  int
}

type Student4 struct {
	Person4
	school bool
}

func (p Person4) sayHello1() {
	fmt.Println("Hello, I am", p.Name)
}

type human interface {
	sayHello1()
}

func bar1(h human) {
	switch h.(type) {
	case Person4:
		fmt.Println("I was pass into bar", h.(Person4).Name)
	case Student4:
		fmt.Println("I was pass into bar", h.(Student4).Name)
	}
}

func main() {
	sa1 := Student4{Person4{"Tom", 18}, true}
	fmt.Println(sa1)

	sa1.sayHello1()

	bar1(sa1)
}
