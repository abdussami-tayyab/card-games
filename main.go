package main

import (
	"github.com/gin-gonic/gin"

	h "github.com/abdussami-tayyab/card-games/api/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/decks", h.CreateDeck)
	router.GET("/decks", h.OpenDeck)

	router.Run()
}
