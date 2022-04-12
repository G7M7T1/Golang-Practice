package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type newPerson struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	jsons := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(jsons)

	var people []newPerson

	err := json.Unmarshal([]byte(jsons), &people)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(people)

	for i, p := range people {
		fmt.Println("Person", i)

		fmt.Println(p.First, p.Last, p.Age)

		for _, val := range p.Sayings {
			fmt.Println(val)
		}

	}
}
