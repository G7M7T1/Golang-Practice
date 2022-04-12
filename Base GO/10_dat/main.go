package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	v, ok := m["Jamess"]
	fmt.Println(v, ok)

	k, done := m["James"]
	fmt.Println(k, done)
}
