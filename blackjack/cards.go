package blackjack

func parseCard(card string) int {
	var value int
	switch card {
		case "ace":
			value = 11
		case "two":
			value = 2
		case "three":
			value = 3
		case "four":
			value = 4
		case "five":
			value = 5
		case "six":
			value = 6
		case "seven":
			value = 7
		case "eight":
			value = 8
		case "nine":
			value = 9
		case "ten", "jack", "queen", "king":
			value = 10
		default:
			value = 0
	}
	return value
}
