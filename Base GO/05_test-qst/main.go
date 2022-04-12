package main

import "fmt"

var x = 1
var y = "hello"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
