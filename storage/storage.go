package storage

import (
	"github.com/abdussami-tayyab/card-games/model"
	"github.com/google/uuid"
)

// AllDecks is a global map to store decks by their ID.
var AllDecks = make(map[uuid.UUID]model.Deck)

// AddDeck adds a new deck to the AllDecks map.
func AddDeck(deck model.Deck) {
	AllDecks[deck.DeckID] = deck
}

// GetDeck retrieves a deck by its ID from the AllDecks map.
func GetDeck(id uuid.UUID) (model.Deck, bool) {
	deck, found := AllDecks[id]
	return deck, found
}
