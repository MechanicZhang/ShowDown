package main

func main() {
	game := NewShowDown(NewDeck(), []Player{NewHumanPlayer(), NewHumanPlayer(), NewHumanPlayer(), NewHumanPlayer()})
	game.Start()
}