package main

type Hand struct {
	cards []Card
}

func NewHand() *Hand {
	return &Hand{make([]Card, 0)}
}

func (h *Hand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) Size() int {
	return len(h.cards)
}