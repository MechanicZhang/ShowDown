package main

import "fmt"

type ExchangeHands struct {
	timer     int
	exchanger IPlayer
	exchangee IPlayer
}

func NewExchangeHands(exchanger IPlayer, exchangee IPlayer) *ExchangeHands {
	e := &ExchangeHands{3, exchanger, exchangee}
	e.exchange()
	return e
}

func (e *ExchangeHands) GetExchanger() IPlayer {
	return e.exchanger
}

func (e *ExchangeHands) GetExchangee() IPlayer {
	return e.exchangee
}

func (e *ExchangeHands) GetCount() int {
	return e.timer
}

func (e *ExchangeHands) exchange() {
	e.printExchange()
	hand := e.exchanger.GetHand()
	e.exchanger.SetHand(e.exchangee.GetHand())
	e.exchangee.SetHand(hand)
}

func (e *ExchangeHands) countdown() {
	e.timer--
	if e.timer == 0 {
		e.exchange()
	}
}

func (e *ExchangeHands) printExchange() {
	fmt.Printf("玩家: %v 與玩家: %v 交換手牌\n", e.exchanger.GetName(), e.exchangee.GetName())
}