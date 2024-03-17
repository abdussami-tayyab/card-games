package services_test

import (
	"testing"

	"github.com/abdussami-tayyab/card-games/services"
	"github.com/google/uuid"
)

func TestCreateDeckWithAllCardsShuffled(t *testing.T) {
	ds := services.DeckService{}
	shuffled := true
	wantedCards := ""

	deck, err := ds.CreateDeck(shuffled, wantedCards)

	if err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	if deck.DeckID == uuid.Nil {
		t.Errorf("Expected deck ID to be generated, but got nil")
	}

	if deck.Shuffled != shuffled {
		t.Errorf("Expected deck to be shuffled, but it was not")
	}

	if len(deck.Cards) != 52 {
		t.Errorf("Expected a full deck to have 52 cards, but got %d", len(deck.Cards))
	}
}

func TestCreateDeckWithWantedCardsShuffled(t *testing.T) {
	ds := services.DeckService{}
	shuffled := true
	wantedCards := "AS,KD,AC,2C,KH"

	deck, err := ds.CreateDeck(shuffled, wantedCards)

	if err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	if deck.DeckID == uuid.Nil {
		t.Errorf("Expected deck ID to be generated, but got nil")
	}

	if deck.Shuffled != shuffled {
		t.Errorf("Expected deck to be shuffled, but it was not")
	}

	if len(deck.Cards) != 5 {
		t.Errorf("Expected deck to have 5 cards, but got %d", len(deck.Cards))
	}

	if deck.Cards[0].Code == "AS" && deck.Cards[1].Code == "KD" && deck.Cards[4].Code == "KH" {
		t.Errorf("Expected deck to match wanted cards, but AS == %v, KD == %v KH == %v", deck.Cards[0].Code, deck.Cards[1].Code, deck.Cards[4].Code)
	}
}

func TestCreateDeckWithAllCardsUnshuffled(t *testing.T) {
	ds := services.DeckService{}
	shuffled := false
	wantedCards := ""

	deck, err := ds.CreateDeck(shuffled, wantedCards)

	if err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	if deck.DeckID == uuid.Nil {
		t.Errorf("Expected deck ID to be generated, but got nil")
	}

	if deck.Shuffled != shuffled {
		t.Errorf("Expected deck to be unshuffled, but it was shuffled")
	}

	if len(deck.Cards) != 52 {
		t.Errorf("Expected a full deck to have 52 cards, but got %d", len(deck.Cards))
	}
}

func TestCreateDeckWithWantedCardsUnshuffled(t *testing.T) {
	ds := services.DeckService{}
	shuffled := false
	wantedCards := "AS,KD,AC,2C,KH"

	deck, err := ds.CreateDeck(shuffled, wantedCards)

	if err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	if deck.DeckID == uuid.Nil {
		t.Errorf("Expected deck ID to be generated, but got nil")
	}

	if deck.Shuffled != shuffled {
		t.Errorf("Expected deck to be unshuffled, but it was shuffled")
	}

	if len(deck.Cards) != 5 {
		t.Errorf("Expected deck to have 5 cards, but got %d", len(deck.Cards))
	}

	if deck.Cards[0].Code != "AS" && deck.Cards[4].Code != "KH" {
		t.Errorf("Expected deck to match wanted cards, AS != %v, KH != %v", deck.Cards[0].Code, deck.Cards[4].Code)
	}
}
