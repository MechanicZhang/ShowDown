package main

import "fmt"

type Human struct {
	*PlayerBase
}

func NewHumanPlayer() *Human {
	return &Human{&PlayerBase{hand: NewHand()}}
}

func (h *Human) TakeTurn() Action {
	fmt.Printf("輪到玩家 %s 的回合, 您要採取什麼行動呢？\n", h.GetName())
	action := Action{player: h, exchangeHands: h.MakeExchangeHandsDecision(), card: h.ShowCard()}
	if eh := action.GetExchangeHands(); eh != nil {
		fmt.Printf("玩家: %v 已使用「交換手牌」特權，與玩家: %v 交換手牌\n", h.GetName(), action.GetExchangeHands().GetExchangee().GetName())
		h.SetExchangeHands(eh)
	}
	return action
}

func (h *Human) ShowCard() Card {
	index, err := promptCardIndex(h.GetHand())
	if err != nil {
		panic(err)
	}

	return h.hand.ShowCard(index - 1)
}

func (h *Human) MakeExchangeHandsDecision() *ExchangeHands {
	if h.hasUsedExchangeHandsDecision {
		return nil
	}
	shouldExchange, err := promptYesNo(ExchangeHandsPrompt)
	if err != nil {
		panic(err)
	}
	if shouldExchange {
		h.hasUsedExchangeHandsDecision = true
		players := filterPlayers(h.showdown.players, func(p IPlayer) bool {
			return h.GetName() != p.GetName()
		})

		targetIndex, err := promptPlayerIndex(players)
		if err != nil {
			panic(err)
		}
		return NewExchangeHands(h, players[targetIndex-1])
	}

	return nil
}
