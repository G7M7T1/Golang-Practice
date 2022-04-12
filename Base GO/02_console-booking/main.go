package main

import "fmt"

func main() {
	var conferenceName = "Go App"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Total tickets:", conferenceTickets, "And", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var UserTickets int

	fmt.Println("Please enter your user name:")
	fmt.Scan(&userName)
	fmt.Println("Please enter your tickets:")
	fmt.Scan(&UserTickets)

	remainingTickets -= UserTickets

	fmt.Println("Your Name:", userName)
	fmt.Println("Your Ticket:", UserTickets)
	fmt.Println("Total tickets:", conferenceTickets, "And", remainingTickets, "are still available")
}
