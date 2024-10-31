package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicket := userTicket > 0 && uint(userTicket) <= remainingTicket
	return isValidName, isValidEmail, isValidTicket
}