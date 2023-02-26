package main

type Action struct {
	player IPlayer
	card   Card
	exchangeHands *ExchangeHands
}

func NewAction(player IPlayer, card Card, exchangeHands *ExchangeHands) Action {
	return Action{
		player: nil,
		card: Card{
			Suit: 0,
			Rank: 0,
		},
		exchangeHands: exchangeHands,
	}
}


func (a *Action) GetPlayer() string {
	return a.player.GetName()
}

func (a *Action) GetCard() Card {
	return a.card
}

func (a *Action) GetExchangeHands() *ExchangeHands {
	return a.exchangeHands
}
