package main

import "fmt"

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
