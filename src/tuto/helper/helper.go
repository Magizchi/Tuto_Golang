package helper

import "strings"

func ValidateUserInpt(firstName string, lastName string, email string, userTickets int, remaningTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickerNumber := userTickets > 0 && userTickets < int(remaningTickets)

	return isValidName, isValidEmail, isValidTickerNumber
}
