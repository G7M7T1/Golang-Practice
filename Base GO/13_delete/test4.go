package main

import "fmt"

func main() {
	xs1 := []string{"a", "b", "c", "d"}
	xs2 := []string{"e", "f", "g", "h"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xs := [][]string{xs1, xs2}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println("i:v", i, ":", v)
	}
}
