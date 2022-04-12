package main

import "fmt"

func main() {
	// 定义一个key:value的哈希表
	m := make(map[string]int)
	// 赋值
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("m:", m)
	// 给变量赋值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// 哈希长度
	fmt.Println("len(m):", len(m))

	// 删除一个哈希值
	delete(m, "k2")
	fmt.Println("m:", m)

	// 定义+初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
