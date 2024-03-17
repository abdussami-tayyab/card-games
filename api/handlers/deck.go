package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/abdussami-tayyab/card-games/model"
	"github.com/abdussami-tayyab/card-games/services"
)

var AllDecks = make(map[uuid.UUID]model.Deck)

// Create a new deck of cards
func CreateDeck(c *gin.Context) {
	shuffle := c.Query("shuffle") == "true"
	wantedCards := c.Query("cards")

	// Create deck using service
	deckService := services.NewDeckService()
	newDeck, err := deckService.CreateDeck(shuffle, wantedCards)
	if err != nil {
		log.Println("Error while creating deck", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create deck"})

		return
	}
	AllDecks[newDeck.DeckID] = newDeck

	response := model.DeckResponse{
		DeckID:    newDeck.DeckID.String(),
		Shuffled:  newDeck.Shuffled,
		Remaining: len(newDeck.Cards),
		Cards:     newDeck.Cards,
	}

	c.IndentedJSON(http.StatusCreated, response)
}

// Open a deck of cards
func OpenDeck(c *gin.Context) {
	deckId := c.Param("id")

	// No UUID passed
	if deckId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UUID found in request"})
		return
	}

	// Invalid UUID passed
	parsedUuid, err := uuid.Parse(deckId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse Deck ID"})
		return
	}

	// Find deck with given uuid
	deck, found := AllDecks[parsedUuid]
	if !found {
		log.Printf("No deck found for UUID %s", deckId)

		c.JSON(http.StatusNotFound, gin.H{"error": "No deck found."})
		return
	}

	response := model.DeckResponse{
		DeckID:    deck.DeckID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	}

	c.IndentedJSON(http.StatusOK, response)
}
