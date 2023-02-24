package main

import "fmt"

type ShowDown struct {
	deck    *Deck
	players []*Player
}

func NewShowDown(deck *Deck, players []*Player) *ShowDown {
	return &ShowDown{
		deck:    deck,
		players: players,
	}
}

func (sd *ShowDown) Start() {
	fmt.Println("deck's size: ", len(sd.deck.cards))
	for i := 0; i < len(sd.deck.cards); i++ {
		fmt.Printf(" %s", sd.deck.cards[i])
	}
	// naming
	// shuffle
	// draw
	// play 13 rounds
	// show winner
}
