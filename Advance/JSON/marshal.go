package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{First: "Jams", Last: "Morry", Age: 18}
	p2 := person{First: "Toms", Last: "Kary", Age: 20}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
