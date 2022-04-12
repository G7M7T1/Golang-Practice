package main

import "fmt"

// 声明user结构体
type user struct {
	name     string
	password string
	age      int
}

func main() {
	// 实例化 user 结构体 user1
	user1 := user{name: "koma", password: "12345678", age: 25}
	fmt.Println(user1.name, user1.password, user1.age)

	// 声明一个指向 user1 的指针 user1p
	user1p := &user1
	fmt.Println(user1p.name, user1p.password, user1p.age)

	// 利用指针给 user1 赋值，会改变 user1 的内容
	user1p.name = "mike"
	user1p.password = "iloveu"
	user1p.age = 20
	fmt.Println(user1.name, user1.password, user1.age)
	fmt.Println(user1p.name, user1p.password, user1p.age)
}
