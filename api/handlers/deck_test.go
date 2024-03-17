package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdussami-tayyab/card-games/api/handlers"
	"github.com/abdussami-tayyab/card-games/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/decks", handlers.CreateDeck)
	router.GET("/decks/:id", handlers.OpenDeck)
	router.POST("/decks/:id/draw", handlers.DrawCards)
	return router
}

func TestCreateDeckWithAllCardsShuffled(t *testing.T) {
	router := SetUpRouter()

	w := performRequest(router, "POST", "/decks?shuffle=true", nil)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if !response.Shuffled {
		t.Errorf("Expected deck to be shuffled, but it was not")
	}

	if response.Remaining != 52 {
		t.Errorf("Expected a full deck to have 52 cards, but got %d", response.Remaining)
	}
}

func TestCreateDeckWithAllCardsUnshuffledAndPartial(t *testing.T) {
	router := SetUpRouter()

	w := performRequest(router, "POST", "/decks?shuffle=false", nil)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response.Shuffled {
		t.Errorf("Expected deck to be un-shuffled, but it was")
	}

	if response.Remaining != 52 {
		t.Errorf("Expected a full deck to have 52 cards, but got %d", response.Remaining)
	}
}

func TestCreateDeckWithAllCardsShuffledAndPartial(t *testing.T) {
	router := SetUpRouter()

	w := performRequest(router, "POST", "/decks?shuffle=true&cards=AS,KD,AC,2C,KH", nil)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if !response.Shuffled {
		t.Errorf("Expected deck to be shuffled, but it was not")
	}

	if response.Remaining != 5 {
		t.Errorf("Expected this partial deck to have 5 cards, but got %d", response.Remaining)
	}
}

func TestCreateDeckWithAllCardsUnshuffled(t *testing.T) {
	router := SetUpRouter()

	w := performRequest(router, "POST", "/decks?shuffle=false&cards=AS,KD,AC,2C,KH", nil)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response.Shuffled {
		t.Errorf("Expected deck to be unshuffled, but it was shuffled")
	}

	if response.Remaining != 5 {
		t.Errorf("Expected this partial deck to have 5 cards, but got %d", response.Remaining)
	}
}

func TestOpenDeckWithValidDeck(t *testing.T) {
	router := SetUpRouter()

	// First create a deck
	w := performRequest(router, "POST", "/decks", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := deckResponse.DeckID

	// Then open a deck via the deck's ID
	w = performRequest(router, "GET", fmt.Sprintf("/decks/%s", deckId), nil)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var openResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &openResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if deckResponse.DeckID != openResponse.DeckID {
		t.Errorf("Both decks don't match %s != %s", deckResponse.DeckID, openResponse.DeckID)
	}

	if deckResponse.Remaining != openResponse.Remaining {
		t.Errorf("Expected the number of cards to match in both decks")
	}

	if deckResponse.Remaining == openResponse.Remaining {
		if len(deckResponse.Cards) > 0 && len(openResponse.Cards) > 0 {
			for i := 0; i < deckResponse.Remaining; i++ {
				if deckResponse.Cards[i].Code != openResponse.Cards[i].Code {
					t.Errorf("Expected all cards in the deck to match the original deck when opened")
					break
				}
			}
		}
	}
}

func TestOpenDeckWithInvalidDeck(t *testing.T) {
	router := SetUpRouter()

	// First create a deck
	w := performRequest(router, "POST", "/decks", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := uuid.NewString()

	// Then open a deck via the deck's ID
	w = performRequest(router, "GET", fmt.Sprintf("/decks/%s", deckId), nil)
	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}

	// Todo: add a check to see if the JSON error response is also correct
}

func TestDrawCardsWithValidDeckAndValidCount(t *testing.T) {
	router := SetUpRouter()

	// First create a deck
	w := performRequest(router, "POST", "/decks", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := deckResponse.DeckID
	total := 52
	count := 3

	// Then draw cards from deck
	w = performRequest(router, "POST", fmt.Sprintf("/decks/%s/draw?count=%d", deckId, count), nil)
	if w.Code != http.StatusOK {
		log.Printf("%v", w.Body)
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
	var drawnCards []model.Card
	if err := json.Unmarshal(w.Body.Bytes(), &drawnCards); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check if count and drawn cards are equal
	if len(drawnCards) != count {
		t.Errorf("Expected 3 cards to be drawn, %d drawn", len(drawnCards))
	}

	// Check if now remaining cards match the right number
	w = performRequest(router, "GET", fmt.Sprintf("/decks/%s", deckId), nil)
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if total-len(drawnCards) != deckResponse.Remaining {
		t.Errorf("Expected deck's cards and drawn cards to match")
	}
}

func TestDrawCardsWithValidDeckAndInvalidCount(t *testing.T) {
	router := SetUpRouter()

	// First create a deck
	w := performRequest(router, "POST", "/decks", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := deckResponse.DeckID

	// Then draw cards from deck
	w = performRequest(router, "POST", fmt.Sprintf("/decks/%s/draw?count=53", deckId), nil)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestDrawCardsWithInvalidDeck(t *testing.T) {
	router := SetUpRouter()

	// First create a deck
	w := performRequest(router, "POST", "/decks", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := uuid.NewString()

	// Then open a deck via the deck's ID
	w = performRequest(router, "GET", fmt.Sprintf("/decks/%s", deckId), nil)
	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}

	// Todo: add a check to see if the JSON error response is also correct
}

func TestDrawCardsFromEmptyDeck(t *testing.T) {
	router := SetUpRouter()

	// First create a deck with one card
	w := performRequest(router, "POST", "/decks?cards=AS", nil)
	var deckResponse model.DeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &deckResponse); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	deckId := deckResponse.DeckID

	// Then draw one card from deck
	w = performRequest(router, "POST", fmt.Sprintf("/decks/%s/draw?count=1", deckId), nil)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Then draw one card again from deck (now empty)
	w = performRequest(router, "POST", fmt.Sprintf("/decks/%s/draw?count=1", deckId), nil)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}
