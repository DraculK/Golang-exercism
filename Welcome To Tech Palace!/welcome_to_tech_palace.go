package techpalace

import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customerUpper := strings.ToUpper(customer)
	welcome := "Welcome to the Tech Palace, " + customerUpper
    return welcome
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := ""
    finalString := ""
	for i:=0; i<numStarsPerLine; i++{
        stars += "*" 
    }
	finalString = stars + "\n" + welcomeMsg + "\n" + stars
    return finalString
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    removeStars := ""
    removeSpaces := ""
	removeBreak := ""
	finalString := ""
	removeStars = strings.ReplaceAll(oldMsg, "*", "")
    removeSpaces = strings.Trim(removeStars, " ")
	removeBreak = strings.Trim(removeSpaces, "\n ")
    finalString = removeBreak
    return finalString
}