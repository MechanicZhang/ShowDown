package main

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

func (h *Human) TakeTurn() Turn {
	turn := Turn{player: h, card: h.hand.cards[0]}
	return turn
}

func (h *Human) AddHandCard(c Card) {
	h.hand.AddCard(c)
}

func (h *Human) ShowCard(c Card) {
	panic("not implemented") // TODO: Implement
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
