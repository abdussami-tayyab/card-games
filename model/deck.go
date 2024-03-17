package model

import (
	"fmt"
	"math/rand"

	"github.com/abdussami-tayyab/card-games/util"
	"github.com/google/uuid"
)

// A deck of cards.
type Deck struct {
	DeckID   uuid.UUID `json:"deck_id"`
	Shuffled bool      `json:"shuffled"`
	Cards    []Card    `json:"cards"`
}

// A type for Deck to respond in handlers.
type DeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards,omitempty"` // as cards are optional in some responses
}

// A Card that belongs to a deck.
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
	Code  string `json:"code"`
}

// Constructor for Card
func NewCard(value, suit, code string) Card {
	card := Card{
		Suit:  suit,
		Value: value,
		Code:  code,
	}

	// Set Code for the card based on value and suit
	if code == "" {
		card.Code = fmt.Sprintf("%c%c", value[0], suit[0])
	}

	// Set Value and Suit for code-only card
	if value == "" && suit == "" && code != "" {
		valueRune := rune(code[0])
		suitRune := rune(code[1])

		card.Value = util.ConvertValue(valueRune)
		card.Suit = util.ConvertSuit(suitRune)
	}

	return card
}

// Shuffle the cards
func Shuffle(cards []Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
