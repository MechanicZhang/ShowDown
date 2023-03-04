package main

type Human struct {
	*PlayerBase
}

func NewHumanPlayer() *Human {
	return &Human{&PlayerBase{hand: NewHand()}}
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
