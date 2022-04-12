package main

import "fmt"

func main() {
	x := [5]int{2, 3, 5, 8, 4}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
