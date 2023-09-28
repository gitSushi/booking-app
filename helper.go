//	̶M̶o̶v̶e̶d̶ ̶t̶o̶ ̶i̶t̶s̶ ̶o̶w̶n̶ ̶p̶a̶c̶k̶a̶g̶e̶
//
// Own packages are in own folders
package main

import "strings"



func validateUserInput(firstname string, lastname string, email string, userTickets int) (bool, bool, bool){
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidUserTicket
}