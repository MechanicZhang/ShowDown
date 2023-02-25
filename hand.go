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

func (h *Hand) GetCards() []Card {
	return h.cards
}

func (h *Hand) ShowCard(index int) Card{
	card := h.cards[index]
	h.cards = append(h.cards[:index], h.cards[index+1:]...)
	return card
}

func (h *Hand) Size() int {
	return len(h.cards)
}