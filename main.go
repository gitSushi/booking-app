package main

import "fmt"

func main() {
	conferenceName := "our conference" // syntactic sugar for implicit variables
	const conferenceTickets uint8 = 50 // There cannot be a negative total ticket
	var remainingTickets uint16 = 50

	// fmt.Printf("type of conferenceName is %T\n", conferenceName)       // string
	// fmt.Printf("type of conferenceTickets is %T\n", conferenceTickets) // int
	// fmt.Printf("type of remainingTickets is %T\n", remainingTickets)   // int

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets still available out of the %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend")

	fmt.Println(conferenceName)
	fmt.Println(&conferenceName) // pointer -> displays memory

	// arrays hold only one type and definite length
	// var bookings [50]string
	// slices are based on arrays
	// alternative syntax
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	var firstname string
	var lastname string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstname)
	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets to book: ")
	fmt.Scan(&userTickets)

	remainingTickets -= uint16(userTickets)
	// assign to an array index
	// bookings[0] = firstname + " " + lastname
	// append to a slice
	bookings = append(bookings, firstname+" "+lastname)

	// fmt.Printf("The whole array or slice: %v\n", bookings)
	// fmt.Printf("First value: %v\n", bookings[0])
	// fmt.Printf("Array or slice type: %T\n", bookings)
	// fmt.Printf("Array or slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("List of all the bookings:\n%v\n", bookings)
}
