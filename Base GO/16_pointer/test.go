package main

import "fmt"

type persons struct {
	name string
}

func main() {
	p1 := persons{name: "Jams"}
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *persons) {
	p.name = "Miss Money"
	// or  choose one of
	(*p).name = "Other Name"
}
