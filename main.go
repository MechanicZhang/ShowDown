package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	game := NewShowDown(NewDeck(),
		[]Player{
			NewHumanPlayer(),
			NewAIPlayer(),
			NewHumanPlayer(),
			NewAIPlayer(),
		},
	)
	game.Start()
}
