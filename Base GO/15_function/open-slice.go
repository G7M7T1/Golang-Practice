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
