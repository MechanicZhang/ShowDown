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

func (ai *AI) TakeTurn() Action {
	fmt.Printf("輪到玩家 %s 的回合, 您要採取什麼行動呢？\n", ai.GetName())
	action := Action{player: ai, exchangeHands: ai.MakeExchangeHandsDecision(), card: ai.ShowCard()}
	if eh := action.GetExchangeHands(); eh != nil {
		ai.SetExchangeHands(eh)
	}
	return action
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
