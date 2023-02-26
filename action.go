package main

type Action struct {
	player Player
	card   Card
}

func (a *Action) GetPlayer() string {
	return a.player.GetName()
}

func (a *Action) GetCard() Card {
	return a.card
}
