package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 8, 6, 4, 2, 3, 5, 8, 7}
	sort.Ints(s)
	fmt.Println(s)
}
