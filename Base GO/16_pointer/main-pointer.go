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
