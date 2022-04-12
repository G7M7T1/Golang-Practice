package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46}

	x = append(x, 47)
	fmt.Println(x)
}
