package main

import (
	"fmt"
	"math/rand"
)

type AI struct {
	*PlayerBase
}

func NewAIPlayer() *AI {
	return &AI{&PlayerBase{hand: NewHand()}}
}

func (ai *AI) ShowCard() Card {
	fmt.Printf("玩家: %v 已自動出牌\n", ai.GetName())
	return ai.hand.ShowCard(rand.Intn(ai.HandSize()))
}

func (ai *AI) MakeExchangeHandsDecision() *ExchangeHands {
	if ai.hasUsedExchangeHandsDecision {
		return nil
	}
	if rand.Intn(2) == 1 {
		ai.hasUsedExchangeHandsDecision = true
		players := filterPlayers(ai.showdown.players, func(p IPlayer) bool {
			return ai.GetName() != p.GetName()
		})
		return NewExchangeHands(ai, players[rand.Intn(len(players))])
	}

	return nil
}
