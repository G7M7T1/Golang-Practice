package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
	}

	fmt.Println(m)
	fmt.Println(m["bond_james"])
	fmt.Println(m["bond_james"][0])

	for k, v := range m {
		for i, v2 := range v {
			fmt.Println("---", k, i, v2)
		}
	}
}
