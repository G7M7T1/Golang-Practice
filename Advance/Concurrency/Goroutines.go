package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func sayHelo(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Helo", name, ":", i)
	}
}

// 主函数
func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	// 同步执行函数
	sayHelo("koma")
	// 异步执行函数
	go sayHelo("xiaoma")
	go sayHelo("iphone")
	go sayHelo("ipad")
	go sayHelo("swiftui")
	// 匿名函数，异步执行
	go func(msg string) {
		fmt.Println("this is a", msg)
	}("lesson")
	// 等待一秒
	time.Sleep(time.Second)
}
