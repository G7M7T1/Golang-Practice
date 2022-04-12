package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	delete(m, "James")
	fmt.Println(m)
}