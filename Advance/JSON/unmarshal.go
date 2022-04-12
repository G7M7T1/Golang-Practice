package main

import (
	"encoding/json"
	"fmt"
)

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Jams","Last":"Morry","Age":18},{"First":"Toms","Last":"Kary","Age":20}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person2{}
	var people []person2
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v)
	}
}
