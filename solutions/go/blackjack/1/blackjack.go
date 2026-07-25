package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val := 0
	switch card {
	case "ace":
		val = 11
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	case "ten", "jack", "queen", "king":
		val = 10
	default:
		val = 0
	}

	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	decision := "H"

	switch {
	case playerSum == 22:
		decision = "P"
	case playerSum == 21:
		if ParseCard(dealerCard) < 10 {
			decision = "W"
		} else {
			decision = "S"
		}
	case playerSum >= 17 && playerSum <= 20:
		decision = "S"
	case (playerSum >= 12 && playerSum <= 16) && ParseCard(dealerCard) < 7:
		decision = "S"
	default:
		decision = "H"
	}

	return decision
}
