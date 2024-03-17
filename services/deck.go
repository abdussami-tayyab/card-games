package services

import (
	"strings"

	"github.com/abdussami-tayyab/card-games/model"
	"github.com/google/uuid"
)

// DeckService provides operations on decks of cards
type DeckService struct{}

// NewDeckService creates a new instance of DeckService
func NewDeckService() *DeckService {
	return &DeckService{}
}

// CreateDeck creates a new deck of cards, optionally shuffling them.
func (ds *DeckService) CreateDeck(shuffled bool, wantedCards string) (model.Deck, error) {
	var cards []model.Card

	if wantedCards == "" {
		cards = getAllCards()
	} else {
		cards = getWantedCards(wantedCards)
	}

	if shuffled {
		model.Shuffle(cards)
	}

	deck := model.Deck{
		DeckID:   uuid.New(),
		Shuffled: shuffled,
		Cards:    cards,
	}

	return deck, nil
}

// Get the full deck, all cards.
func getAllCards() (cards []model.Card) {
	values := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "TEN", "JACK", "QUEEN", "KING"}
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, model.NewCard(value, suit, ""))
		}
	}

	return
}

// Get only the cards wanted by client.
func getWantedCards(wantedCards string) (cards []model.Card) {
	wantedCardsArr := strings.Split(wantedCards, ",")

	for _, code := range wantedCardsArr {
		cards = append(cards, model.NewCard("", "", code))
	}

	return
}
