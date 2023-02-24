package main

func main() {
	game := NewShowDown(NewDeck(), []*Player{})
	game.Start()
}