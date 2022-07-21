package techpalace

import "strings"

// message for (most) customers scanning their membership cards
func welcomeMessage(cust string) string {
	var genericMessage string = "Welcome to the Tech Palace, "
	var user string = strings.ToUpper(cust)
	return genericMessage + user
}

// message for loyal customers scanning their membership cards
func addBorder(msg string, stars int) string {
	var line1 string = strings.Repeat("*", stars)
	var line2 string = msg
	var line3 string = strings.Repeat("*", stars)

	completeMessage := line1 + "\n" + line2 + "\n" + line3
	return completeMessage
}

// removes * and trims leading and trailing spaces from string
func cleanupMessage(msg string) string {
	var noBorder string = strings.ReplaceAll(msg, "*", "")
	var textOnly string = strings.TrimSpace(noBorder)
	return textOnly
}
