package main

import (
	"fmt"
	"math/rand"
	"time"
)

type AI struct {
	name  string
	point int
	hand  *Hand
}

func NewAIPlayer() *AI {
	return &AI{hand: NewHand()}
}

func (ai *AI) SetName(name string) {
	ai.name = name
}

func (ai *AI) GetName() string {
	return ai.name
}

func (ai *AI) TakeTurn() Turn {
	fmt.Printf("輪到玩家 %s 的回合, 已自動出牌。\n", ai.name)
	randIndex := rand.Intn(ai.HandSize()) + 1
	card := ai.hand.ShowCard(randIndex - 1)
	turn := Turn{player: ai, card: card}
	time.Sleep(500 * time.Millisecond)
	return turn
}

func (ai *AI) AddHandCard(c Card) {
	ai.hand.AddCard(c)
}

func (ai *AI) ShowCard(index int) Card {
	return ai.hand.ShowCard(index)
}

func (ai *AI) GainPoint() {
	ai.point++
}

func (ai *AI) GetPoint() int {
	return ai.point
}

func (ai *AI) MakeExchangeHandsDecision() {
	panic("not implemented") // TODO: Implement
}

func (ai *AI) HandSize() int {
	return ai.hand.Size()
}
