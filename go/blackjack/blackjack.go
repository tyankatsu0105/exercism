package blackjack

const (
	Ace   = 11
	Two   = 2
	Three = 3
	Four  = 4
	Five  = 5
	Six   = 6
	Seven = 7
	Eight = 8
	Nine  = 9
	Ten   = 10
	Jack  = 10
	Queen = 10
	King  = 10
	Other = 0
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := Other

	switch card {
	case "ace":
		value = Ace
	case "two":
		value = Two
	case "three":
		value = Three
	case "four":
		value = Four
	case "five":
		value = Five
	case "six":
		value = Six
	case "seven":
		value = Seven
	case "eight":
		value = Eight
	case "nine":
		value = Nine
	case "ten":
		value = Ten
	case "jack":
		value = Jack
	case "queen":
		value = Queen
	case "king":
		value = King
	}

	return value
}

const (
	Stand             = "S"
	Hit               = "H"
	Split             = "P"
	Automatically_win = "W"
)

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardValue1 := ParseCard(card1)
	cardValue2 := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	result := Automatically_win

	switch true {
	case cardValue1 == Ace && cardValue2 == Ace:
		result = Split
	case cardValue1+cardValue2 == 21:
		if dealerCardValue == Ace || dealerCardValue == Ten || dealerCardValue == Jack || dealerCardValue == Queen || dealerCardValue == King {
			result = Stand
		} else {
			result = Automatically_win
		}
	case cardValue1+cardValue2 <= 11:
		result = Hit
	case 17 <= cardValue1+cardValue2 && cardValue1+cardValue2 <= 20:
		result = Stand
	case 12 <= cardValue1+cardValue2 && cardValue1+cardValue2 <= 16:
		if Seven <= dealerCardValue {
			result = Hit
		} else {
			result = Stand
		}
	}

	return result
}
