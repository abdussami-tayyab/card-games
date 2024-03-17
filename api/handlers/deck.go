package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	}

	c.IndentedJSON(http.StatusCreated, response)
}

// Open a deck of cards
func OpenDeck(c *gin.Context) {
	deckId := c.Param("id")

	// Find deck with given uuid
	status, response := ValidateDeck(deckId)
	if status != http.StatusOK {
		c.JSON(status, response)
		return
	}

	parsedUuid, _ := uuid.Parse(deckId)
	deck := AllDecks[parsedUuid]

	deckResponse := model.DeckResponse{
		DeckID:    deck.DeckID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	}

	c.IndentedJSON(http.StatusOK, deckResponse)
}

// Draw cards from deck
func DrawCards(c *gin.Context) {
	deckId := c.Param("id")
	countStr := c.Query("count")

	// Find deck with given uuid
	status, response := ValidateDeck(deckId)
	if status != http.StatusOK {
		c.JSON(status, response)
		return
	}

	// All clear, fetch values
	parsedUuid, _ := uuid.Parse(deckId)
	deck := AllDecks[parsedUuid]
	cards := deck.Cards

	// Count validation
	// count checks
	// 1. count is invalid
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'count' parameter, must be an integer"})
		return
	}

	// 2. user requests <0 cards
	if count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must draw at least 1 card"})
		return
	}

	// 3. user requests >remaining cards
	if count > len(cards) {
		errorMsg := fmt.Sprintf("Deck has %d cards but you are requesting %d", len(cards), count)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	// Draw cards from deck
	startIndex := len(cards) - count
	drawn := cards[startIndex:]

	// Update original deck
	deck.Cards = cards[:startIndex]
	AllDecks[parsedUuid] = deck

	clientResponse := map[string][]model.Card{}
	clientResponse["cards"] = drawn

	c.IndentedJSON(http.StatusOK, clientResponse)
}

// Validate if a deck passes all requirements
// Todo: Can be moved to a middleware
func ValidateDeck(deckId string) (status int, response map[string]interface{}) {
	// No UUID passed
	if deckId == "" {
		return http.StatusBadRequest, gin.H{"error": "No UUID found in request"}
	}

	// Invalid UUID passed
	parsedUuid, err := uuid.Parse(deckId)
	if err != nil {
		return http.StatusBadRequest, gin.H{"error": "Unable to parse Malformed Deck ID"}
	}

	// Find deck with given uuid
	_, found := AllDecks[parsedUuid]
	if !found {
		log.Printf("No deck found for UUID %s", deckId)
		return http.StatusNotFound, gin.H{"error": "No deck found."}
	}

	return http.StatusOK, nil
}
