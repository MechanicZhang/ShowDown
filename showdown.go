package main

import "fmt"

type ShowDown struct {
	deck    *Deck
	players []Player
}

type Turn struct {
	player Player
	card   Card
}

func (t *Turn) GetPlayer() string {
	return t.player.GetName()
}

func (t *Turn) GetCard() Card {
	return t.card
}

func NewShowDown(deck *Deck, players []Player) *ShowDown {
	return &ShowDown{
		deck:    deck,
		players: players,
	}
}

func (sd *ShowDown) Start() {
	// naming
	for i := 0; i < len(sd.players); i++ {
		sd.players[i].SetName(fmt.Sprintf("Player %d", i+1))
	}
	sd.greeting()
	// shuffle
	sd.deck.Shuffle()
	// draw
	for i := 0; i < 13; i++ {
		for j := 0; j < len(sd.players); j++ {
			sd.players[j].AddHandCard(sd.deck.DrawCard())
		}
	}
	fmt.Println("After drawing...")
	for i := 0; i < len(sd.players); i++ {
		fmt.Printf(" %v has %v cards\n", sd.players[i].GetName(), sd.players[i].HandSize())
	}
	// play 13 rounds
	for i := 0; i < 13; i++ {
		turns := make([]Turn, len(sd.players))
		for j := 0; j < len(sd.players); j++ {
			turn := sd.players[j].TakeTurn()
			turns[j] = turn 
		}

		for j := 0; j < len(turns); j++ {
			fmt.Printf("%v's turn: %v\n", turns[j].GetPlayer(), turns[j].GetCard())
		}

		// compare turns
		winTurn := turns[0]
		for j := 1; j < len(turns); j++ {
			if turns[j].GetCard().ShowDown(winTurn.GetCard()) > 0 {
				winTurn = turns[j]
			}
		}

		fmt.Println("Winner is", winTurn.GetPlayer())
		// give point
		winTurn.player.GainPoint()
		// show card
	}
	// show winner
	sd.gameover()
}

func (sd *ShowDown) greeting() {
	fmt.Println("Welcome to the ShowDown!")
	fmt.Println("Players:")
	for i := 0; i < len(sd.players); i++ {
		fmt.Printf(" %s", sd.players[i].GetName())
	}
	fmt.Println()
}

func (sd *ShowDown) gameover() {
	fmt.Println("Game over!")
	winner := sd.players[0]
	for i := 1; i < len(sd.players); i++ {
		if sd.players[i].GetPoint() > winner.GetPoint() {
			winner = sd.players[i]
		}
	}
	fmt.Printf("The winner is %v, who got %v points\n", winner.GetName(), winner.GetPoint())
}