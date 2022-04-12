package main

import "fmt"

func main() {
	defer fun1()
	fun2()
}

func fun1() {
	fmt.Printf("fun1\n")
}

func fun2() {
	fmt.Printf("fun2\n")
}
