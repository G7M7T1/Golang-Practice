package main

import "fmt"

func main() {
	s1 := foo1()
	fmt.Println(s1)

	x := bar2()
	i := x()
	fmt.Println(i)
}

func foo1() string {
	s := "Hello there"
	return s
}

func bar2() func() int {
	return func() int {
		return 123
	}
}
