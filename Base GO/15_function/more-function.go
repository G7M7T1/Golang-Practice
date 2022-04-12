package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum: %d\n", sum)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
