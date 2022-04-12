package main

import "fmt"

func main() {
	x := [5]int{2, 3, 5, 8, 4}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
