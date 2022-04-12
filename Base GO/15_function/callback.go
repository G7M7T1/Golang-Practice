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
