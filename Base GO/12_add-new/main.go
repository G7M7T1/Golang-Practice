package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}

	m["Tara"] = 27

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
