package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
        case "ace":
    		return 11
    	case "two":
    		return 2
        case "three":
    		return 3
        case "four":
    		return 4
        case "five":
    		return 5
        case "six":
    		return 6
        case "seven":
    		return 7
        case "eight":
    		return 8
        case "nine":
    		return 9
        case "ten":
    		return 10
        case "jack":
    		return 10
        case "queen":
    		return 10
        case "king":
    		return 10
        case "other":
    		return 0
    }
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
    c2 := ParseCard(card2)
    c3 := ParseCard(dealerCard)
	cards_sum := c1 + c2
    switch {
        case c1 == 11 && c2 == 11:
    		return "P"
        case cards_sum == 21 && c3 <= 9:
    		return "W"
        case cards_sum == 21 && c3 >= 10:
    		return "S"
        case cards_sum >= 17 && cards_sum <= 21:
        	return "S"
        case cards_sum >= 12 && cards_sum <= 16 && c3 >= 7:
        	return "H"
        case cards_sum >= 12 && cards_sum <= 16 && c3 <= 7:
        	return "S"
        case cards_sum <= 11:
    		return "H"
    }
	return "0"
}
