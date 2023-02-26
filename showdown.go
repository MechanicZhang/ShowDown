package main

import "fmt"

type ShowDown struct {
	deck    *Deck
	players []Player
	actions []Action
}

func NewShowDown(deck *Deck, players []Player) *ShowDown {
	return &ShowDown{
		deck:    deck,
		players: players,
		actions: make([]Action, len(players)),
	}
}

func (sd *ShowDown) Start() {
	sd.naming()
	sd.greeting()
	sd.deck.Shuffle()
	sd.drawing()
	sd.playing()
	sd.gameover()
}

func (sd *ShowDown) naming() {
	for i := 0; i < len(sd.players); i++ {
		fmt.Printf("請輸入第 %v 位玩家名稱: ", i+1)
		var name string
		fmt.Scanln(&name)
		sd.players[i].SetName(name)
	}
}

func (sd *ShowDown) greeting() {
	fmt.Println("---------------------------------")
	fmt.Println("歡迎來到比大小遊戲！")
	fmt.Println("參賽玩家:")
	for i := 0; i < len(sd.players); i++ {
		fmt.Printf(" %s", sd.players[i].GetName())
	}
	fmt.Printf("\n\n")
}

func (sd *ShowDown) drawing() {
	fmt.Println("所有玩家抽牌中...")
	for i := 0; i < 13; i++ {
		for j := 0; j < len(sd.players); j++ {
			sd.players[j].AddHandCard(sd.deck.DrawCard())
		}
	}
}

func (sd *ShowDown) playing() {
	for i := 0; i < 13; i++ {
		fmt.Println("---------------------------------")
		fmt.Printf("開始第 %v 回合\n", i+1)
		for j := 0; j < len(sd.players); j++ {
			fmt.Printf("輪到玩家 %s 的回合, 您要採取什麼行動呢？\n", sd.players[j].GetName())
			action := sd.players[j].TakeTurn()
			sd.actions[j] = action
		}

		for j := 0; j < len(sd.actions); j++ {
			fmt.Printf("玩家 %v 出了 %v\n", sd.actions[j].GetPlayer(), sd.actions[j].GetCard())
		}

		// compare turns
		winTurn := sd.actions[0]
		for j := 1; j < len(sd.actions); j++ {
			if sd.actions[j].GetCard().ShowDown(winTurn.GetCard()) > 0 {
				winTurn = sd.actions[j]
			}
		}

		fmt.Printf("本回合最大的牌為 %v，玩家 %v 獲勝。\n\n", winTurn.GetCard(), winTurn.GetPlayer())
		// give point
		winTurn.player.GainPoint()
		// show card
	}
}

func (sd *ShowDown) gameover() {
	fmt.Println("遊戲結束!")
	winner := sd.players[0]
	for i := 1; i < len(sd.players); i++ {
		if sd.players[i].GetPoint() > winner.GetPoint() {
			winner = sd.players[i]
		}
	}
	fmt.Printf("本局遊戲獲勝的玩家是 %v, 一共贏得了 %v 分。\n", winner.GetName(), winner.GetPoint())
}
