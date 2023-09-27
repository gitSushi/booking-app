package main

import (
	"fmt"
	"strings"
)

// Package Level Variables
// Seems to be like global scope variables
// They cannot de defined as implicit variables
const conferenceTickets uint8 = 50
var conferenceName = "our conference"
var remainingTickets uint16 = 50
var bookings []string

func main() {

	// Place as package level variables
	/*
	conferenceName := "our conference" // syntactic sugar for implicit variables
	const conferenceTickets uint8 = 50 // There cannot be a negative total ticket
	var remainingTickets uint16 = 50
	*/

	// fmt.Printf("type of conferenceName is %T\n", conferenceName)       // string
	// fmt.Printf("type of conferenceTickets is %T\n", conferenceTickets) // int
	// fmt.Printf("type of remainingTickets is %T\n", remainingTickets)   // int

	greetVisitor() // Replaces following commented block
	/*
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets still available out of the %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend")
	*/

	fmt.Println(conferenceName)
	fmt.Println(&conferenceName) // pointer -> displays memory

	// Place as package level variable
	/*
	// arrays hold only one type and definite length
	// var bookings [50]string
	// slices are based on arrays
	// alternative syntax
	var bookings []string // slice of strings
	// var bookings = []string{}
	// bookings := []string{}
	*/

	// Conditions in a for loop
	// for remainingTickets > 0 && len(bookings) < 50 {
	// Implicit condition true
	for {
		firstname, lastname, email, userTickets := getUserInput() // Replaces following commented block
		/*
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
		*/

		isValidName, isValidEmail, isValidUserTicket := validateUserInput(firstname, lastname, email, userTickets) // Replaces following commented block
		/*
		isValidName := len(firstname) >= 2 && len(lastname) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTicket := userTickets > 0 && userTickets <= int(remainingTickets)
		*/

		// Shows the use of continue keyword
		// if userTickets > int(remainingTickets) {
		// 	fmt.Printf("We only have %v remaining tickets. You cannot book %v tickets.\n", remainingTickets, userTickets)
		// 	continue
		// }

		// if userTickets <= int(remainingTickets) {
		if isValidName && isValidEmail && isValidUserTicket {
			bookings = bookTicket(userTickets, firstname, lastname, email) // Replaces following commented block
			/*
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
			*/
		
			fmt.Printf("List of all the bookings:\n%v\n", bookings)
			
			// Replaces following commented block
			fmt.Printf("List of all the firstnames of bookings:\n%v\n", getFirstnames())
			/*
			firstnames := []string{}
			
			// Blank identifier
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			
			for index, booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, strconv.Itoa(index)+" "+names[0])
			}

			fmt.Printf("List of all the firstnames of bookings:\n%v\n", firstnames)
			*/
	
	
			if remainingTickets == 0 {
				// end program
				fmt.Println("The conference is booked out. Come back next year.")
				break
			}
		} else {
			// fmt.Printf("We only have %v remaining tickets. You cannot book %v tickets.\n", remainingTickets, userTickets)
			// fmt.Println("Your input data is invalid. Try again.")
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is missing the '@' sign.")
			}
			if !isValidUserTicket {
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}


	showSwitchExample()
}

// Function accessing package level variable
// func greetVisitor(conferenceName string, remainingTickets uint16, conferenceTickets uint8){
func greetVisitor(){
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets still available out of the %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend")
}

// func getFirstnames(bookings []string) []string {
func getFirstnames() []string {
	firstnames := []string{}

	// Blank identifier
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}

	return firstnames
}

// func validateUserInput(firstname string, lastname string, email string, userTickets int, remainingTickets uint16) (bool, bool, bool){
func validateUserInput(firstname string, lastname string, email string, userTickets int) (bool, bool, bool){
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidUserTicket
}

func getUserInput() (string, string, string, int){
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

	return firstname, lastname, email, userTickets
}

// func bookTicket(remainingTickets uint16, userTickets int, bookings []string, firstname string, lastname string, email string, conferenceName string) ([]string) {
func bookTicket(userTickets int, firstname string, lastname string, email string) ([]string) {
	remainingTickets -= uint16(userTickets)
	bookings = append(bookings, firstname+" "+lastname)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}

func showSwitchExample(){
	city := "Istanbul"

	switch city {
		case "New York", "Washington":
			// Code for New York
		case "Singapore":
			// Code for Singapore
		case "London":
			// Code for London
		default:
			res := "No valid city were selected."
			returnedBytes := len(res)
			fmt.Printf("Res length is %v == the number of bytes returned\n", returnedBytes)
			byteReturnedByFmt, err := fmt.Printf("%v\n", res)
			fmt.Printf("%v is the number of bytes returned and error %v\n", byteReturnedByFmt, err)
			if returnedBytes != byteReturnedByFmt {
				fmt.Println("For some unknown reason, they are not equal !?!")
			}
	}
}