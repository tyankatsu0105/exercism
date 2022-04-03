package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	// return "Welcome to my party, " + name + "!"
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// return "Happy birthday " + name + "! " + "You are now " + fmt.Sprint(age) + " years old!"
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// return Welcome(name) + "\n" + 
	// "You have been assigned to table " + fmt.Sprintf("%03d", table) + ". " + "Your table is " + direction + ", " + "exactly " + fmt.Sprintf("%.1f", distance) +  " meters from here." + "\n" + 
	// "You will be sitting next to " + neighbor + "."

	result := ""

	result += fmt.Sprintf("Welcome to my party, %s!", name)
	result += "\n"
	result += fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	result += "\n"
	result += fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return result
}
