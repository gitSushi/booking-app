package helper

import "strings"

// Using PascalCase let's you export a function
// It also works for variables
func ValidateUserInput(firstname string, lastname string, email string, userTickets int, remainingTickets uint16) (bool, bool, bool){
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidUserTicket
}