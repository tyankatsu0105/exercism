package techpalace

import (
	"strings"
)

const BORDER_CHAR = "*"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {	
	border := strings.Repeat(BORDER_CHAR, numStarsPerLine)

	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	borderRemovedMessage := strings.Replace(oldMsg, BORDER_CHAR, " ", -1)

	return strings.TrimSpace(borderRemovedMessage)
}
