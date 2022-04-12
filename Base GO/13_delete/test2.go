package main

import "fmt"

func main() {
	x := [5]int{42, 43, 44, 45, 46}

	fmt.Println(x[1:3])
	fmt.Println(x[:3])
	fmt.Println(x[2:])
}
