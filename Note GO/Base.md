# Base **[Go Programming Language](https://go.dev/)**

## **`const` and** **`var`**

```go
const conferenceTickets = 50
var remainingTickets = 50
```

## `Println`

```go
fmt.Println("Hello World")
```

## `Printf` and `%v`

```go
var conferenceName = "Your Name"
fmt.Printf("Welcome %v to GOlang\n", Name)
```

## `dataType` and `%T` and `:=`

```go
var userName string = "Name"
var userMoney int = 60
```

```go
var userMoney int = 60
fmt.Printf("userMoney %T", userMoney)
```

```go
userMoney := 60
fmt.Println(userMoney)
// 自动分配
```

## `Scan`

```go
var userName string
var UserMoney int

fmt.Scan(&userName)
fmt.Scan(&UserMoney)
	
fmt.Println("userName", userName)
fmt.Println("UserMoney", UserMoney)
```

## `type`

```go
type NewType int
var b NewType = 3
fmt.Println(b)
```

## `convert type`

```go
type NewType int

func main() {
	var a int
	var b NewType = 3

	a = 10
	fmt.Println(a)

	a = int(b)
	fmt.Println(a)

	fmt.Println(b)
	fmt.Printf("Type: %T\n", b)
}
```

## `Sprintf`

```go
var x int = 1
var y string = "hello"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
```

## `iota`

```go
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
```

## `switch` and `fallthrough`

```go
package main

import "fmt"

func main() {

	var a = 10

	switch {

	case a == 10:
		fmt.Println("This should print")
		fallthrough

	case a == 11:
		fmt.Println("This should not print")
		fallthrough

	case a == 12:
		fmt.Println("This should print")
		fallthrough

	case (3 == 4):
		fmt.Println("This should not print")
		fallthrough

	case (4 == 4):
		fmt.Println("This should print 4 4")

	default:
		fmt.Println("This is default case")
	}
}
```

```go
This should print
This should not print
This should print
This should not print
This should print 4 4
```

## `Array`

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
}
```

```go
func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
}
```

## `range`

```go
package main

import "fmt"

func main() {
	x := [5]int{2, 3, 5, 8, 4}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
```

```go
0 2
1 3
2 5
3 8
4 4
```

## `slice` and `:`

```go
func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[2:])
	fmt.Println(x[2:3])
}
```

## `append` and `...` and `make`

```go
func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	x = append(x, 6, 7, 8)
	fmt.Println(x)
	
	y := []int{9, 10, 11}
	x = append(x, y...)
	fmt.Println(x)
	
}
```

```go
func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

//  [0 0 0 0 0]
//  5
//  10
```

## `[][]` `muti array`

```go
package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
```

```go
[James Bond Chocolate Martini]
[Miss Moneypenny Strawberry Hazelnut]
[[James Bond Chocolate Martini] [Miss Moneypenny Strawberry Hazelnut]]
```

## `map`

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
}
```

## `ok`

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	v, ok := m["Jamess"]
	fmt.Println(v, ok)

	k, ok := m["James"]
	fmt.Println(k, ok)

}
```

```go
0 false
32 true
```

## `if ok`

```go
package main

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	if v, ok := m["James"]; ok {
		println(v)
	}
}
```

```go
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
}
```

```go
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
```

```go
32
map[James:32 Miss Moneypenny:27 Tara:27]
James 32
Miss Moneypenny 27
Tara 27
```

## `delete`

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	delete(m, "James")
	fmt.Println(m)
}
```

## `[][]string`

```go
package main

import "fmt"

func main() {
	xs1 := []string{"a", "b", "c", "d"}
	xs2 := []string{"e", "f", "g", "h"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xs := [][]string{xs1, xs2}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println("i:v", i, ":", v)
	}
}
```

```go
[a b c d]
[e f g h]
[[a b c d] [e f g h]]
i:v     0 : [a b c d]
i:v     1 : [e f g h]
```

## `map[string][]string`

```go
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
```

```go
map[bond_james:[Shaken, not stirred Martinis] moneypenny_miss:[James Bond Literature Computer Science]]
[Shaken, not stirred Martinis]
Shaken, not stirred
--- bond_james 0 Shaken, not stirred
--- bond_james 1 Martinis
--- moneypenny_miss 0 James Bond
--- moneypenny_miss 1 Literature
--- moneypenny_miss 2 Computer Science
```

## `struct`

```go
package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
}
```

or

```go
package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
}
```

## `x**,**y` or `x**,**y**,**z`

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println(x, y, z)

	x, y, z = 1, 2, 3
	fmt.Println(x, y, z)
}
```

```go
0 0 0
1 2 3
```

## `struct` and `[]`

```go
package main

import "fmt"

type person5 struct {
	name string
	age  int
	like []string
}

func main() {
	p1 := person5{
		name: "John",
		age:  25,
		like: []string{"coding", "running"},
	}
	fmt.Println(p1)

	for i, v := range p1.like {
		fmt.Println(i, v)
	}
}
```

```go
{John 25 [coding running]}
0 coding
1 running
```

## `struct with struct`

```go
package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}
```

```go
{{2 red} true}
{{4 blue} true}
```

## `Function`

```go
package main

import "fmt"

func main() {
	z := add(1, 2)
	fmt.Println(z)

	sayHello("Hi")
}

func add(x int, y int) int {
	return x + y
}

func sayHello(s string) {
	fmt.Println("Hello", s)
}
```

```go
3
Hello Hi
```

## `function return`

```go
package main

import "fmt"

func main() {
	z := add(1, 2)
	fmt.Println(z)

	sayHello("Hi")

	x, y := moused("Hello", "World")
	fmt.Println(x)
	fmt.Println(y)
}

func add(x int, y int) int {
	return x + y
}

func sayHello(s string) {
	fmt.Println("Hello", s)
}

func moused(fn string, ln string) (string, bool) {
	aa := fmt.Sprint("Hello ", fn, ln)
	bb := true
	return aa, bb
}
```

```go
3
Hello Hi
Hello HelloWorld
true
```

## `... function`

```go
package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum: %d\n", sum)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
```

or 

```go
package main

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum1(xi...)
	println(x)
}

func sum1(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
```

## `defer`

```go
package main

import "fmt"

func main() {
	defer fun1()
	fun2()
}

func fun1() {
	fmt.Printf("fun1\n")
}

func fun2() {
	fmt.Printf("fun2\n")
}
```

```go
fun2
fun1
```

## `method`

```go
package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

type Student3 struct {
	Person3
	school bool
}

func (p Person3) sayHello() {
	fmt.Println("Hello, I am", p.Name)
}

func main() {
	sa1 := Student3{Person3{"Tom", 18}, true}
	fmt.Println(sa1)

	sa1.sayHello()
}
```

```go
{{Tom 18} true}
Hello, I am Tom
```

## `log`

```go
package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString", myString)
}
```

## `interface` `(important)`

```go
package main

import "fmt"

type Person4 struct {
	Name string
	Age  int
}

type Student4 struct {
	Person4
	school bool
}

func (p Person4) sayHello1() {
	fmt.Println("Hello, I am", p.Name)
}

type human interface {
	sayHello1()
}

func bar1(h human) {
	switch h.(type) {
	case Person4:
		fmt.Println("I was pass into bar", h.(Person4).Name)
	case Student4:
		fmt.Println("I was pass into bar", h.(Student4).Name)
	}
}

func main() {
	sa1 := Student4{Person4{"Tom", 18}, true}
	fmt.Println(sa1)

	sa1.sayHello1()

	bar1(sa1)
}
```

```go
{{Tom 18} true}
Hello, I am Tom
I was pass into bar Tom
```

## `anonymous function`

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello There")
	}()
}
```

or

```go
package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello There")
	}

	f()
}
```

## `return`

```go
package main

import "fmt"

func main() {
	s1 := foo1()
	fmt.Println(s1)
}

func foo1() string {
	s := "Hello there"
	return s
}
```

## `Return Function`

```go
package main

import "fmt"

func main() {
	s1 := foo1()
	fmt.Println(s1)

	x := bar2()
	i := x()
	fmt.Println(i)
}

func foo1() string {
	s := "Hello there"
	return s
}

func bar2() func() int {
	return func() int {
		return 123
	}
}
```

```go
123
```

## `CallBack Function`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")

	totalNumbers := sum3(3, 5, 4, 7, 8, 9)
	fmt.Println(totalNumbers)

	numberList1 := []int{1, 2, 5, 6, 3, 2, 4, 5, 8}
	newNumbers3 := evenNumber(sum3, numberList1...)
	fmt.Println(newNumbers3)
}

func sum3(xi ...int) int {
	total := 0

	for _, i := range xi {
		total += i
	}
	return total
}

func evenNumber(f func(xi ...int) int, vi ...int) int {
	var newYi []int

	for _, v := range vi {
		if v%2 == 0 {
			newYi = append(newYi, v)
		}
	}
	fmt.Println(newYi)
	return f(newYi...)
}
```

```go
Hello
36
[2 6 2 4 8]
22
```

## `quset function`

```go
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{radius: 12.5}
	squa := square{length: 5}

	info(circ)
	info(squa)
}
```

```go
490.8738521234052
25
```

# `Pointer` or `指针`

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
