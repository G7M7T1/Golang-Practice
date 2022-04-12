package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user1 struct {
	Name string
	Age  int
}

func main() {
	u1 := user1{"James", 18}
	u2 := user1{"Tomas", 22}

	users := []user1{u1, u2}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
