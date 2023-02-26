package main

import "fmt"

type Human struct {
	name  string
	point int
	hand  *Hand
}

func NewHumanPlayer() *Human {
	return &Human{hand: NewHand()}
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) TakeTurn() Action {
	fmt.Printf("   ")
	for _, card := range h.GetHandCards() {
		fmt.Printf(" %v", card)
	}
	fmt.Println()
	var index int
	fmt.Scanf("%d", &index)
	for index < 1 || index > h.hand.Size() {
		size := h.hand.Size()
		fmt.Printf("輸入錯誤！您目前有 %v 張牌，請輸入介於 1 ~ %v 的正整數: ", size, size)
		fmt.Scanf("%d", &index)
	}
	card := h.hand.ShowCard(index - 1)
	turn := Action{player: h, card: card}
	return turn
}

func (h *Human) GetHandCards() []Card {
	return h.hand.GetCards()
}

func (h *Human) AddHandCard(c Card) {
	h.hand.AddCard(c)
}

func (h *Human) ShowCard(index int) Card {
	return h.hand.ShowCard(index)
}

func (h *Human) GainPoint() {
	h.point++
}

func (h *Human) GetPoint() int {
	return h.point
}

func (h *Human) MakeExchangeHandsDecision() {
	panic("not implemented") // TODO: Implement
}

func (h *Human) HandSize() int {
	return h.hand.Size()
}
