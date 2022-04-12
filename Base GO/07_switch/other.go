package main

import (
	"fmt"
	"time"
)

func main() {

	player_level := 1
	switch player_level {
	case 1:
		fmt.Println("aa，bb，cc")
	case 2:
		fmt.Println("dd，ee，ff")
	case 3:
		fmt.Println("gg，hh，jj")
	case 4, 5, 6:
		fmt.Println("kk")
	default:
		fmt.Println("ll")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("ee，rr")
	default:
		fmt.Println("qq")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("ww：ee")
	default:
		fmt.Println("ww：ff")
	}
}
