package util

// Mappings for card value.
var valueMappings = map[rune]string{
	'A': "ACE",
	'2': "2",
	'3': "3",
	'4': "4",
	'5': "5",
	'6': "6",
	'7': "7",
	'8': "8",
	'9': "9",
	'T': "TEN",
	'J': "JACK",
	'Q': "QUEEN",
	'K': "KING",
}

// Mappings for suits.
var suitMappings = map[rune]string{
	'S': "SPADES",
	'D': "DIAMONDS",
	'C': "CLUBS",
	'H': "HEARTS",
}

// ConvertValue converts a single character to its full value name.
func ConvertValue(char rune) string {
	if val, ok := valueMappings[char]; ok {
		return val
	}

	return string(char)
}

// ConvertSuit converts a single character to its full suit name.
func ConvertSuit(char rune) string {
	if val, ok := suitMappings[char]; ok {
		return val
	}

	return string(char)
}
