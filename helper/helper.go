package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {

	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isvalidName, isvalidEmail, isvalidTicketNumber

}
 