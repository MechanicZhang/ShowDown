package main

type Player interface {
	SetName(name string)
	GetName() string
	TakeTurn() Turn
	AddHandCard(c Card)
	ShowCard(index int) Card
	GainPoint()
	GetPoint() int
	MakeExchangeHandsDecision()

	HandSize() int
}