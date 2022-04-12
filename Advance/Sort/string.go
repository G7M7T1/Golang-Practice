package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"rr", "dd", "ubc", "ast", "aaf"}
	sort.Strings(s)
	fmt.Println(s)
}
