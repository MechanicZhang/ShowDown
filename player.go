package main

type IPlayer interface {
	// common methods
	GetName() string
	SetName(name string)
	SetShowDown(showdown *ShowDown)
	GetPoint() int
	GainPoint()
	GetHand() []Card
	SetHand(cards []Card)
	AddHandCard(c Card)
	HandSize() int
	GetExchangeHands() *ExchangeHands
	SetExchangeHands(eh *ExchangeHands)

	// abstract methods
	TakeTurn() Action
	ShowCard() Card
	MakeExchangeHandsDecision() *ExchangeHands
}

type PlayerBase struct {
	name          string
	point         int
	hand          *Hand
	showdown      *ShowDown
	exchangeHands *ExchangeHands

	hasUsedExchangeHandsDecision bool
}

func (pb *PlayerBase) GetName() string {
	return pb.name
}

func (pb *PlayerBase) SetName(name string) {
	pb.name = name
}

func (pb *PlayerBase) SetShowDown(showdown *ShowDown) {
	pb.showdown = showdown
}

func (pb *PlayerBase) GetPoint() int {
	return pb.point
}

func (pb *PlayerBase) GainPoint() {
	pb.point++
}

func (pb *PlayerBase) GetHand() []Card {
	return pb.hand.GetCards()
}

func (pb *PlayerBase) SetHand(cards []Card) {
	pb.hand.SetCards(cards)
}

func (pb *PlayerBase) AddHandCard(c Card) {
	pb.hand.AddCard(c)
}

func (pb *PlayerBase) HandSize() int {
	return pb.hand.Size()
}

func (pb *PlayerBase) GetExchangeHands() *ExchangeHands {
	return pb.exchangeHands
}

func (pb *PlayerBase) SetExchangeHands(eh *ExchangeHands) {
	pb.exchangeHands = eh
}

// helper functions
func filterPlayers(players []IPlayer, filter func(IPlayer) bool) []IPlayer {
	filtered := make([]IPlayer, 0)
	for _, v := range players {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
