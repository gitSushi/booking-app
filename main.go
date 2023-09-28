// A package is a collection of Go file(s)
package main

/*
import (
	"booking-app/helper"
	"fmt"
	"strings"
	"strconv"
)
*/
import "fmt"

//  ̶M̶o̶v̶e̶d̶ ̶t̶o̶ ̶h̶e̶l̶p̶e̶r̶/̶h̶e̶l̶p̶e̶r̶.̶g̶o̶
/*
// var conferenceName = "our conference"
*/

// Package Level Variables == package scope variables
//  ̶S̶e̶e̶m̶s̶ ̶t̶o̶ ̶b̶e̶ ̶l̶i̶k̶e̶ ̶g̶l̶o̶b̶a̶l̶ ̶s̶c̶o̶p̶e̶ ̶v̶a̶r̶i̶a̶b̶l̶e̶s̶
// They cannot de defined as implicit variables
const conferenceTickets uint8 = 50
var conferenceName = "our conference"
var remainingTickets uint16 = 50
var bookings = make([]UserStruct, 0)
// Changed for struct
/*
// To create a slice with make we need to define an initial size
var bookings = make([]map[string]string, 0)
*/
// Changed for maps
/*
var bookings []string
*/

// Creating a new type
type UserStruct struct {
	firstname string
	lastname string
	email string
	userTickets	int
}

func main() {

	// Placed as package level variables
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

	// Placed as package level variable
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
		
			// NOTE :
			// With map, it prints each key in alphabetical order
			// With struct, it prints in defined order
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

// Function with parameters to access local level variables
// func greetVisitor(conferenceName string, remainingTickets uint16, conferenceTickets uint8){
// Function accessing package level variables
func greetVisitor(){
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets still available out of the %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend")
}

// func getFirstnames(bookings []string) []string {
func getFirstnames() []string {
	// Local scope variable within getFirstnames function
	firstnames := []string{}
	
	// Blank identifier
	for _, booking := range bookings {
		// Changed because of maps
		/*
		// Local scope variable within the for loop
		// Is NOT known outside of it. Not even inside the getFirstnames function
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
		*/
		// Changed because of struct
		/*
		firstnames = append(firstnames, booking["firstname"])
		*/
		firstnames = append(firstnames, booking.firstname)
	}

	return firstnames
}

// Moved to helper.go
/*
// func validateUserInput(firstname string, lastname string, email string, userTickets int, remainingTickets uint16) (bool, bool, bool){
func validateUserInput(firstname string, lastname string, email string, userTickets int) (bool, bool, bool){
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidUserTicket
}
*/

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
// func bookTicket(userTickets int, firstname string, lastname string, email string) ([]string) {
// func bookTicket(userTickets int, firstname string, lastname string, email string) ([]map[string]string) {
func bookTicket(userTickets int, firstname string, lastname string, email string) ([]UserStruct) {
	remainingTickets -= uint16(userTickets)

	var userData = UserStruct {
		firstname: firstname,
		lastname: lastname,
		email: email,
		userTickets: userTickets,
	}
	// Changed for struct
	/*
	// Create a map for a user
	var userData = make(map[string]string)
	userData["firstname"] = firstname
	userData["lastname"] = lastname
	userData["email"] = email
	// userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10)
	userData["userTickets"] = strconv.Itoa(userTickets)
	*/

	// bookings = append(bookings, firstname+" "+lastname)
	bookings = append(bookings, userData)

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