package main

import (
	"github.com/gin-gonic/gin"

	h "github.com/abdussami-tayyab/card-games/api/handlers"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/decks", h.CreateDeck)
	router.GET("/api/v1/decks/:id", h.OpenDeck)
	router.POST("/api/v1/decks/:id/draw", h.DrawCards)

	router.Run()
}
