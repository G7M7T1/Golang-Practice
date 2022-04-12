# `Advance`

## `Pointer` or `指针`

```go
package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("myString", myString)

	changeUsingPointer(&myString)
	log.Println("myString", myString)
}

func changeUsingPointer(s *string) {
	newValue := "RED"
	*s = newValue
}
```

```go
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
```

## `Json`

Make Json, Read and Write

### `Marshal`

```go
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
```

```go
[{Jams Morry 18} {Toms Kary 20}]
[{"First":"Jams","Last":"Morry","Age":18},{"First":"Toms","Last":"Kary","Age":20}]
```

### `Unmarshal`

```go
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
```

or

```go
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
```

```go
[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]
[{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]
Person 0
James Bond 32
Shaken, not stirred
Youth is no guarantee of innovation
In his majesty's royal service
Person 1
Miss Moneypenny 27
James, it is soo good to see you
Would you like me to take care of that for you, James?
I would really prefer to be a secret agent myself.
Person 2
M Hmmmm 54
Oh, James. You didn't.
Dear God, what has James done now?
Can someone please tell me where James Bond is?
```

## `sort Ints`

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 8, 6, 4, 2, 3, 5, 8, 7}
	sort.Ints(s)
	fmt.Println(s)
}
```

## `sort Strings`

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"rr", "dd", "ubc", "ast", "aaf"}
	sort.Strings(s)
	fmt.Println(s)
}
```