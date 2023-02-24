package main

type Player interface {
	SetName(name string)
	TakeTurn()
	AddHandCard(c Card)
	ShowCard(c Card)
	GainPoint()
	MakeExchangeHandsDecision()
}