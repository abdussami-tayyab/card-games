package main

import (
	"github.com/gin-gonic/gin"

	h "github.com/abdussami-tayyab/card-games/api/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/decks", h.CreateDeck)
	router.GET("/decks/:id", h.OpenDeck)
	router.POST("/decks/:id/draw", h.DrawCards)

	router.Run()
}
