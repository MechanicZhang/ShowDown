package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	players := []IPlayer{
		NewHumanPlayer(), NewAIPlayer(), NewAIPlayer(), NewAIPlayer(),
	}
	game := NewShowDown(NewDeck(), players)
	game.Start()
}
