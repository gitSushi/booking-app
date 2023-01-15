package main

import "fmt"

func main() {
	var conferenceName = "our conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking app")
	fmt.Println("We have", remainingTickets, "tickets still available out of the", conferenceTickets, ".")
	fmt.Println("Get your ticket here to attend")

}
