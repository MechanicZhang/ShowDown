package main

const (
	Club = iota
	Diamond
	Heart
	Spade
)

const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var (
	suits = map[int]string{
		Club:    "♣",
		Diamond: "♦",
		Heart:   "♥",
		Spade:   "♠",
	}

	ranks = map[int]string{
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
		Ten:   "10",
		Jack:  "J",
		Queen: "Q",
		King:  "K",
		Ace:   "A",
	}	
)

type Card struct {
	Suit int
	Rank int
}

func (c Card) String() string {
	return suits[c.Suit] + ranks[c.Rank]
}

func (c Card) ShowDown(c2 Card) int {
	if c.Rank == c2.Rank {
		return c.Suit - c2.Suit
	}
	return c.Rank - c2.Rank
}