package main

import "fmt"

type ShowDown struct {
	deck    *Deck
	players []IPlayer
	actions []Action
}

func NewShowDown(deck *Deck, players []IPlayer) *ShowDown {
	showdown := &ShowDown{
		deck:    deck,
		players: players,
		actions: make([]Action, 0),
	}
	for _, player := range players {
		player.SetShowDown(showdown)
	}
	return showdown
}

func (sd *ShowDown) Start() {
	sd.namePlayers()
	sd.greeting()
	sd.deck.Shuffle()
	sd.drawHands()
	sd.playing()
	sd.gameover()
}

func (sd *ShowDown) namePlayers() {
	for i := 0; i < len(sd.players); i++ {
		name, err := promptPlayerName(fmt.Sprintf(TypePlayerNamePromt, i+1))
		if err != nil {
			panic(err)
		}
		sd.players[i].NameSelf(name)
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

func (sd *ShowDown) drawHands() {
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
			sd.actions = append(sd.actions, takeTurn(sd.players[j]))

		}
		sd.showdown()
		sd.countdown()
		sd.clearActions()
	}
}

func takeTurn(p IPlayer) Action {
	fmt.Printf("輪到玩家 %s 的回合, 您要採取什麼行動呢？\n", p.GetName())
	action := Action{player: p, exchangeHands: p.MakeExchangeHandsDecision(), card: p.ShowCard()}
	if eh := action.GetExchangeHands(); eh != nil {
		p.SetExchangeHands(eh)
	}
	return action
}

func (sd *ShowDown) showdown() {
	printRoundResult(sd.actions)
	winTurn := sd.actions[0]
	for j := 1; j < len(sd.actions); j++ {
		if sd.actions[j].GetCard().ShowDown(winTurn.GetCard()) > 0 {
			winTurn = sd.actions[j]
		}
	}

	fmt.Printf("%v 最大，玩家 %v 獲勝。\n\n", winTurn.GetCard(), winTurn.GetPlayer())
	winTurn.player.GainPoint()
}

func printRoundResult(actions []Action) {
	fmt.Println()
	fmt.Println("本回合的結果:")
	fmt.Printf("     ")
	for i := 0; i < len(actions); i++ {
		fmt.Printf(" %v", actions[i].GetCard())
	}
	fmt.Println()
}

func (sd *ShowDown) countdown() {
	for i := 0; i < len(sd.players); i++ {
		if eh := sd.players[i].GetExchangeHands(); eh != nil {
			eh.countdown()
			if eh.GetCount() > 0 {
				fmt.Printf("玩家 %v 的交換特權剩餘 %v 回合。\n", sd.players[i].GetName(), eh.GetCount())
			}
		}
	}
}

func (sd *ShowDown) clearActions() {
	sd.actions = sd.actions[:0]
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
