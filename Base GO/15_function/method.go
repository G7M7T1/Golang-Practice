package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

type Student3 struct {
	Person3
	school bool
}

func (p Person3) sayHello() {
	fmt.Println("Hello, I am", p.Name)
}

func main() {
	sa1 := Student3{Person3{"Tom", 18}, true}
	fmt.Println(sa1)
	sa1.sayHello()
}
