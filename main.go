package main

import "fmt"

func main() {
	conferenceName := "our conference" // syntactic sugar for implicit variables
	const conferenceTickets uint8 = 50 // There cannot be a negative total ticket
	var remainingTickets uint16 = 50

	fmt.Printf("type of conferenceName is %T\n", conferenceName)       // string
	fmt.Printf("type of conferenceTickets is %T\n", conferenceTickets) // int
	fmt.Printf("type of remainingTickets is %T\n", remainingTickets)   // int

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets still available out of the %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend")

	var username string
	var userTickets int
	// ask user for their name

	username = "Tom"
	userTickets = 2

	fmt.Printf("User named %v, booked %v tickets.\n", username, userTickets)
}
