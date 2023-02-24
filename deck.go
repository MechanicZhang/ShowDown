package main

import "fmt"

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	cards := make([]Card, 52)
	for i := 0; i < 52; i++ {
		cards[i] = Card{Suit: i / 13, Rank: i % 13}
	}
	return &Deck{cards: cards}
}

func (d *Deck) Shuffle() {
	fmt.Println("Shuffling the deck...")
}

func (d *Deck) DrawCard() Card {
	return Card{Suit: 0, Rank: 0}
}
