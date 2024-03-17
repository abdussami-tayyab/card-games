# :black_joker: Decks, Cards, and Games (♠️, ♦️, ♣️, ♥️)

This repository hosts a backend API designed for card games, built using Golang and the Gin web framework.

## API Overview
The API supports the following functionalities:
- **Create a Deck**: Generates a deck of cards. By default, a deck contains 52 cards (13 from each suit: Spades, Diamonds, Clubs, Hearts, in order). Users can specify a subset of cards to include in the deck. The API also supports shuffling the deck.

Example response for creating a deck:
```
{
  "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
  "shuffled": false,
  "remaining": 30
}
```
***
- **Open a Deck**: Retrieves details about a specific deck, including its ID, shuffle status, remaining cards, and the cards themselves.

Example response for opening a deck:
```
{
  "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
  "shuffled": false,
  "remaining": 3,
  "cards": [
    {"value": "ACE", "suit": "SPADES", "code": "AS"},
    {"value": "KING", "suit": "HEARTS", "code": "KH"},
    {"value": "8", "suit": "CLUBS", "code": "8C"}
  ]
}
```
***
- **Draw Cards**: Allows drawing a specified number of cards from a deck.
Example response for drawing cards:
```
{
  "cards": [
    {"value": "QUEEN", "suit": "HEARTS", "code": "QH"},
    {"value": "4", "suit": "DIAMONDS", "code": "4D"}
  ]
}
```

## Getting Started
### Prerequisites
- Ensure you have Go installed on your machine (preferably 1.20+).
- Clone this repository and navigate into the project directory.

### Build the API
Compile the application:
`go build ./...`

### Run the API
Start the API server:
`go run main.go`

### Testing the API
You can test the API by running automated tests by:
`go test ./...`

There is also a Postman Collection

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/33670565-5896f70f-2af9-465e-94aa-97d9eb12100a?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33670565-5896f70f-2af9-465e-94aa-97d9eb12100a%26entityType%3Dcollection%26workspaceId%3D6f7ad2da-6947-4ea7-9a37-6b34d529ad12)

### Support
For questions or feedback, please contact me at tayyab.abdussami@gmail.com.

Thank you for checking out my project! This is my second project in Go, and I'm eager to continue learning and improving.
