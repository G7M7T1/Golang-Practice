package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

const (
	_ = iota // 0
	a        // 1
	b        // 2
	c        // 3
)

func main() {
	fmt.Println(kb, "||", mb, "||", gb)
	fmt.Println(a, "||", b, "||", c)
}
