package main

import (
	"fmt"
	"math/rand"
)

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
	fmt.Println("即將開始新局，洗牌中...")
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) DrawCard() Card {
	top := d.cards[0]
	d.cards = d.cards[1:]
	return top
}
